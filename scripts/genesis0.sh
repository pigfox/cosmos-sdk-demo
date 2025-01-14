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

# Step 3: Remove the existing key if it exists
echo "Removing existing key 'pigfox'..."
simd keys delete pigfox --home "./" --keyring-backend "test" --yes || true  # Use || true to prevent failure if the key does not exist

# Step 4: Add a new genesis account using the test keyring backend
echo "Creating new key 'pigfox'..."
simd keys add pigfox --home "./" --keyring-backend "test"

# Fetch the address of the newly created key
ADDRESS=$(simd keys show pigfox -a --home "./" --keyring-backend "test")
echo "Adding genesis account for address: $ADDRESS"
simd genesis add-genesis-account "$ADDRESS" 100000000stake --home "./"

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

# Step 6: Display the content of the genesis file
echo "Genesis file created and validated successfully."
#cat ./config/genesis.json

#./scripts/copy-generated-genesis.sh