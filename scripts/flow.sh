#!/bin/bash

# Enable error handling
set -e

# Trap errors and print the line number
trap 'echo "Error occurred on line $LINENO"' ERR

# Clear the terminal
clear

# Check if required commands are available
if ! command -v simd >/dev/null 2>&1; then
    echo "Error: 'simd' command not found. Please install the Cosmos SDK."
    exit 1
fi

if ! command -v jq >/dev/null 2>&1; then
    echo "Error: 'jq' command not found. Please install jq for JSON processing."
    exit 1
fi

# Function to convert a string to lowercase
to_lowercase() {
    echo "$1" | tr '[:upper:]' '[:lower:]'
}

# Print simd version
simd version

# Add a validator key if it doesn't exist
if ! simd keys show my-validator >/dev/null 2>&1; then
    echo "Adding validator key..."
    echo "y" | simd keys add my-validator
fi

# Get the validator address
validator_address=$(simd keys show my-validator -a)
if [ -z "$validator_address" ]; then
    echo "Error: Failed to retrieve validator address."
    exit 1
fi
echo "Validator address: $validator_address"

# Ensure the address is in lowercase (to avoid Bech32 error)
validator_address=$(to_lowercase "$validator_address")

# Get the validator public key
validator_pubkey=$(simd tendermint show-validator)
if [ -z "$validator_pubkey" ]; then
    echo "Error: Failed to retrieve validator public key."
    exit 1
fi
echo "Validator public key: $validator_pubkey"

# Use the public key directly as PUBKEY_JSON
PUBKEY_JSON="$validator_pubkey"

# Print the constructed PUBKEY_JSON
echo "Constructed PUBKEY_JSON: $PUBKEY_JSON"

# Convert validator address to operator address
validator_operator_address=$(simd keys parse "$validator_address" | jq -r .valoper_address)
if [ -z "$validator_operator_address" ]; then
    echo "Error: Failed to convert validator address to operator address."
    exit 1
fi
echo "Validator operator address: $validator_operator_address"

# Check and create the genesis file from template if it doesn't exist
if [ ! -f "../genesis_template.json" ]; then
    echo "Error: Genesis template file not found!"
    exit 1
fi

# Create a temporary JSON file and replace required fields
temp_json_file=$(mktemp)

# Use jq to replace placeholders in the genesis template
jq --arg address "$validator_operator_address" \
   --argjson pubkey "$PUBKEY_JSON" \
   '.auth.accounts[0].address = $address | .auth.accounts[0].pubkey = $pubkey' \
   "../genesis_template.json" > "$temp_json_file"

# Validate the JSON
if ! jq empty "$temp_json_file"; then
    echo "Error: JSON is invalid."
    exit 1
fi
echo "JSON is valid."

# Write the updated JSON to genesis.json
genesis_path="./genesis.json"
cp "$temp_json_file" "$genesis_path"
echo "Generated genesis.json written successfully."

# Copy the new genesis file to the simapp config directory
config_path="$HOME/.simapp/config/genesis.json"
cp "$genesis_path" "$config_path" || {
    echo "Error: Failed to copy genesis.json to $config_path."
    exit 1
}

echo "Updated genesis.json copied to $config_path."

# Validate the genesis file
if ! simd validate-genesis; then
    echo "Error: Genesis validation failed. Please check the Bech32 address formatting."
    exit 1
fi

# Output final message
echo "Setup completed successfully. You can now start the chain!"
