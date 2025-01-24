#!/bin/bash
clear

set -x
set -e
trap 'rc=$?; echo "error code $rc at $LINENO"; exit $rc' ERR

# Define costants
HOME_DIR=/home/$USER/.simd
KEYRING_BACKEND=test
CHAIN_ID=mychain

# Remove existing genesis.json file
rm -rf $HOME_DIR/config/genesis.json

# Generate keys
echo "y" | simd keys add mykey --keyring-backend $KEYRING_BACKEND --home $HOME_DIR

# Get the address of the key
my_key_address=$(simd keys show mykey --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --address)

# Get the validator address
my_validator_address=$(simd keys show mykey --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --bech val --address)

# Validate the validator address prefix
if [[ ! "$my_validator_address" =~ ^cosmosvaloper ]]; then
  echo "Error: Invalid validator address prefix. Expected 'cosmosvaloper', got '${my_validator_address}'."
  exit 1
fi

# Add genesis account (using the retrieved address)
simd genesis add-genesis-account $my_key_address 10000stake #--append

# Delegate stake to the validator using the dynamically retrieved validator address
simd tx staking delegate $my_validator_address 500000stake --from mykey --home $HOME_DIR --keyring-backend $KEYRING_BACKEND --chain-id xzxzxz --gas auto --fees 5000stake -y

# Collect genesis transactions
simd collect-gentxs --home $HOME_DIR

# Initialize the chain
simd init mychain --chain-id mychain --home $HOME_DIR

# Start the node
simd start --home $HOME_DIR &

# Wait for the node to start
sleep 5

# Check node status
simd status --home $HOME_DIR
exit
# Delegate stake to the validator 
simd tx staking delegate $my_validator_address 1000000000stake --from mykey --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --chain-id mychain --gas auto --fees 10000stake -y

# Create a simple transaction (e.g., bank send)
simd tx bank send $my_key_address another_address 1000000stake --from mykey --keyring-backend $KEYRING_BACKEND --home $HOME_DIR --chain-id mychain --gas auto --fees 10000stake -y 

# Query account balance
simd query account $my_key_address --home $HOME_DIR --chain-id mychain

# Query validator information
simd query staking validator $my_validator_address --home $HOME_DIR --chain-id mychain

# Stop the node
simd stop --home $HOME_DIR