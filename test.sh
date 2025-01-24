#!/bin/bash
clear

set -x
set -e
trap 'rc=$?; echo "Error code $rc at line $LINENO"; exit $rc' ERR

# Define constants
HOME_DIR="/home/$USER/.simd"
KEYRING_BACKEND="test"
KEY_NAME="mykey"
CHAIN_ID="chainid"
MONIKER="pigfox"
INITIAL_DELEGATION_AMOUNT="1000000000000stake"
GENESIS_ACCOUNT_AMOUNT="1000000000000stake"
RECIPIENT_ADDRESS="cosmos1t0ky8yvljcu0f5zd4rgm02n9h4tehy6m0k6u44"

# Clean previous data
rm -rf "$HOME_DIR"
mkdir -p "$HOME_DIR"

# Print simd version
simd version

# Initialize the chain
simd init "$MONIKER" --chain-id "$CHAIN_ID" --home "$HOME_DIR" --overwrite --log_level debug

# Check for `simd init` errors
if [ $? -ne 0 ]; then
  echo "Error during simd init. Please check the output above for more details."
  exit 1
fi

# Generate keys
if ! simd keys add "$KEY_NAME" --keyring-backend "$KEYRING_BACKEND" --home "$HOME_DIR" <<< "y"; then
  echo "Error: Key generation failed"
  exit 1
fi

# Get the addresses
my_key_address=$(simd keys show "$KEY_NAME" --keyring-backend "$KEYRING_BACKEND" --home "$HOME_DIR" --address)
my_validator_address=$(simd keys show "$KEY_NAME" --keyring-backend "$KEYRING_BACKEND" --home "$HOME_DIR" --bech val --address)

# Validate the validator address prefix
if [[ ! "$my_validator_address" =~ ^cosmosvaloper ]]; then
  echo "Error: Invalid validator address prefix. Expected 'cosmosvaloper', got '${my_validator_address}'."
  exit 1
fi

# Add genesis account (using the retrieved address)
simd genesis add-genesis-account "$my_key_address" "$GENESIS_ACCOUNT_AMOUNT"

# Create a genesis transaction for delegation
simd tx staking delegate $my_validator_address $INITIAL_DELEGATION_AMOUNT \
    --from $KEY_NAME \
    --keyring-backend "$KEYRING_BACKEND" \
    --home "$HOME_DIR" \
    --chain-id "$CHAIN_ID" \
    --generate-only \
    > $HOME_DIR/config/gentx

# Collect genesis transactions
#if ! simd collect-gentxs --home "$HOME_DIR"; then
#  echo "Error: Failed to collect gentxs."
#  exit 1
#fi

# Validate the genesis file
if ! simd validate-genesis --home "$HOME_DIR"; then
  echo "Error: Invalid genesis file."
  exit 1
fi

# Start the node
simd start --home "$HOME_DIR" &

# Wait for the node to start
echo "Waiting for node to start..."
sleep 30

# Check node status
simd status --home "$HOME_DIR"

# Create a simple transaction (e.g., bank send)
simd tx bank send "$my_key_address" "$RECIPIENT_ADDRESS" 1000000stake \
  --from "$KEY_NAME" \
  --keyring-backend "$KEYRING_BACKEND" \
  --home "$HOME_DIR" \
  --chain-id "$CHAIN_ID" \
  --gas auto \
  --fees 10000stake \
  -y

# Query account balance
simd query account "$my_key_address" --home "$HOME_DIR" --chain-id "$CHAIN_ID"

# Query validator information
simd query staking validator "$my_validator_address" --home "$HOME_DIR" --chain-id "$CHAIN_ID"

# Stop the node
pkill -f "simd start --home $HOME_DIR"