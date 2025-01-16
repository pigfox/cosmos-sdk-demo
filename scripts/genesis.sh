#!/bin/bash

# Enable debugging and error handling
set -x
set -eux  # Show commands and exit on any error
clear
KEYRING_BACKEND="test"
CHAIN_ID="1234567890"

# Trap errors and print the line number
trap 'echo "Error occurred on line $LINENO"' ERR

echo "---------Current script name: $0---------"

# Step 0: Clean up any existing configuration and keys
echo "Removing existing configuration and keys..."
rm -rf ./config

# List and delete all existing keys
KEY_LIST=$(simd keys list --home ./ --keyring-backend "$KEYRING_BACKEND" --output json | jq -r '.[].name' || echo "No keys to delete.")
if [[ "$KEY_LIST" != "No keys to delete." ]]; then
  for key in $KEY_LIST; do
    echo "Deleting key: $key"
    simd keys delete "$key" --home ./ --keyring-backend "$KEYRING_BACKEND" -y
  done
else
  echo "No keys to delete."
fi
simd keys list --home ./ --keyring-backend "$KEYRING_BACKEND"

# Step 1: Display the current simd version
echo "Displaying simd version..."
simd version

# Step 2: Initialize the blockchain
echo "Initializing the blockchain..."
simd init "pigfox" --chain-id "$CHAIN_ID" --home "./"

# Step 3: Create a new key and genesis account
KEY_NAME="pigfox_$(date +%s)"
echo "Creating new key: $KEY_NAME"
simd keys add "$KEY_NAME" --home "./" --keyring-backend "$KEYRING_BACKEND"
ADDRESS=$(simd keys show "$KEY_NAME" -a --home "./" --keyring-backend "$KEYRING_BACKEND")
simd genesis add-genesis-account "$ADDRESS" 100000000stake --home "./"

# Step 4: Add validator to genesis
# Fetch the validator address
VALIDATOR_PUBKEY=$(simd tendermint show-validator --home "./")

# Ensure VALIDATOR_PUBKEY is non-empty
if [[ -z "$VALIDATOR_PUBKEY" ]]; then
  echo "Error: Validator public key not found!"
  exit 1
fi
echo "Validator Public Key: $VALIDATOR_PUBKEY"

# Extract the validator address from the public key
VALIDATOR_ADDRESS=$(simd debug pubkey "$VALIDATOR_PUBKEY" | jq -r '.address' || echo "Address extraction failed")

# Ensure the address is non-empty
if [[ -z "$VALIDATOR_ADDRESS" ]]; then
  echo "Error: Validator address not found in pubkey debug output!"
  exit 1
fi

# Convert VALIDATOR_ADDRESS to lowercase
VALIDATOR_ADDRESS=$(echo "$VALIDATOR_ADDRESS" | tr '[:upper:]' '[:lower:]')
echo "Validator address (lowercase): $VALIDATOR_ADDRESS"

# Validate genesis file existence and contents
if [ ! -f "./config/genesis.json" ]; then
  echo "Error: Genesis file not found!"
  exit 1
fi
echo "------------------------------------"
echo "Genesis file validated successfully."
echo "------------------------------------"

# Step 5: Delegate tokens to validator
echo "Delegating tokens to validator..."
echo "Validator Address (before using): $VALIDATOR_ADDRESS"

simd tx staking delegate "$VALIDATOR_ADDRESS" 500000stake \
  --from "$KEY_NAME" --chain-id "$CHAIN_ID" --home "./" \
  --keyring-backend "$KEYRING_BACKEND" --broadcast-mode "sync" --yes \
  --fees 10000stake

# Step 6: Create a second key and transfer tokens
SECOND_KEY_NAME="receiver_$(date +%s)"
echo "Creating second key: $SECOND_KEY_NAME"
simd keys add "$SECOND_KEY_NAME" --home "./" --keyring-backend "$KEYRING_BACKEND"
SECOND_ADDRESS=$(simd keys show "$SECOND_KEY_NAME" -a --home "./" --keyring-backend "$KEYRING_BACKEND")

echo "Transferring tokens to $SECOND_ADDRESS..."
simd tx bank send "$ADDRESS" "$SECOND_ADDRESS" 100000stake \
  --chain-id "$CHAIN_ID" --home "./" --keyring-backend "$KEYRING_BACKEND" --yes

# Step 7: Submit and vote on a governance proposal
echo "Submitting governance proposal..."
simd tx gov submit-proposal text "Increase block size limit" \
  --description "A proposal to increase block size" --deposit 1000000stake \
  --from "$KEY_NAME" --chain-id "$CHAIN_ID" --home "./" --keyring-backend "$KEYRING_BACKEND" --yes

echo "Fetching latest proposal ID..."
PROPOSAL_ID=$(simd query gov proposals --output json | jq -r '.proposals[-1].proposal_id')

echo "Voting on proposal ID $PROPOSAL_ID..."
simd tx gov vote "$PROPOSAL_ID" yes \
  --from "$KEY_NAME" --chain-id "$CHAIN_ID" --home "./" --keyring-backend "$KEYRING_BACKEND" --yes

echo "Blockchain setup and transactions completed successfully!"
