#!/bin/sh

# Enable debugging and error handling
set -x  # Show commands as they execute
set -e  # Exit script on any error
clear
echo "---------Current script name: $0---------"

# Step 0: Clean up any existing configuration
echo "Removing existing configuration..."
rm -rf ./config

# Step 1: Display the current simd version
echo "Displaying simd version..."
simd version

# Step 2: Initialize the blockchain and create the genesis file
echo "Creating genesis file..."
simd init "pigfox" --chain-id "1234567890" --home "./"

# Step 3: Generate a new key with a unique name per execution
KEY_NAME="pigfox_$(date +%s)"
echo "Creating new key '$KEY_NAME'..."
simd keys add "$KEY_NAME" --home "./" --keyring-backend "test"

# Fetch the address and public key of the newly created key
ADDRESS=$(simd keys show "$KEY_NAME" -a --home "./" --keyring-backend "test")
PUB_KEY=$(simd keys show "$KEY_NAME" -p --home "./" --keyring-backend "test" | jq -r '.key')
echo "Adding genesis account for address: $ADDRESS with public key: $PUB_KEY"

# Add the genesis account
simd genesis add-genesis-account "$ADDRESS" 100000000stake --home "./"

# Update the genesis file manually to inject the public key
GENESIS_FILE="./config/genesis.json"
jq --arg address "$ADDRESS" --arg pubkey "$PUB_KEY" \
  '(.app_state.auth.accounts[] | select(.address == $address)).pub_key = {"@type": "/cosmos.crypto.secp256k1.PubKey", "key": $pubkey}' \
  "$GENESIS_FILE" > "${GENESIS_FILE}.tmp" && mv "${GENESIS_FILE}.tmp" "$GENESIS_FILE"

# Step 4: Add the validator address to the genesis file
echo "Fetching validator address..."
VALIDATOR_ADDRESS=$(simd tendermint show-validator --home "./")
echo "Validator address: $VALIDATOR_ADDRESS"

# Step 5: Validate the genesis file
echo "Validating genesis file..."
if [ -f "./config/genesis.json" ]; then
  echo "Genesis file exists. Validating..."
  simd genesis validate --home "./"
  ls -l ./config/genesis.json
else
  echo "Error: Genesis file not found!"
  exit 1
fi

echo "Genesis file created and validated successfully."

#Crete example transactions below

