#!/bin/sh
clear

set -x
set -e

echo "Current user: $USER"

# Check if the node is running and show current keys, starting it if necessary
simd status || (echo "Node is not running, starting the node..." && simd start --home /home/"$USER"/.simapp &)

# List the current keys in the keyring
simd keys list --home /home/"$USER"/.simapp --keyring-backend test --output json

# Show the address of the key 'my-key'
my_key_address=$(simd keys show my-key --keyring-backend test --home /home/"$USER"/.simapp --address)

# Get the validator address for 'my-key'
my_validator_address=$(simd keys show my-key --keyring-backend test --home /home/"$USER"/.simapp --bech val --address)

# Validate the validator address prefix
if [[ ! "$my_validator_address" =~ ^cosmosvaloper ]]; then
  echo "Error: Invalid validator address prefix. Expected 'cosmosvaloper', got '${my_validator_address}'."
  exit 1
fi

# Delete the key 'my-key' if it exists
simd keys delete my-key --keyring-backend test --home /home/"$USER"/.simapp --yes || echo "Key 'my-key' doesn't exist."

# Add the key again
simd keys add my-key --keyring-backend test --home /home/"$USER"/.simapp

# Add genesis account (using the retrieved address)
simd genesis add-genesis-account $my_key_address 10000stake --append

# Wait for the node to be fully operational
echo "Waiting for the node to be fully operational..."
sleep 5 # Adjust sleep time if necessary to ensure the node is ready

# Delegate stake to the validator using the dynamically retrieved validator address
simd tx staking delegate "$my_validator_address" 500000stake --from my-key --home /home/"$USER"/.simapp --keyring-backend test --chain-id xzxzxz --gas auto --fees 5000stake -y

# Collect genesis transactions
simd collect-gentxs --home /home/"$USER"/.simapp

# Initialize the chain
simd init pigfox --chain-id xzxzxz --home /home/"$USER"/.simapp

# Start the node and verify the status
simd start --home /home/"$USER"/.simapp &
simd status

# Ensure the node has fully synced
echo "Ensuring the node is fully synced..."
simd status | grep -q "Syncing: false" || echo "Node is not fully synced, please wait a bit longer."
echo "It works"