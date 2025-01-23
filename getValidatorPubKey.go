package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func getValidatorPubKey(address string) ValidatorKeyData {
	fmt.Println("getValidatorPubKey()")

	// Command to list keys
	validatorPubKeyCmd := []string{
		"keys", "list",
		"--home", settings.AppHomeDir,
		"--keyring-backend", settings.KeyringBackend,
		"--output", "json",
	}

	// Execute the command
	cmd := exec.Command("simd", validatorPubKeyCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\nOutput: %s\n", err, string(output))
		os.Exit(1)
	}

	// Unmarshal the JSON output into a slice of ValidatorKeyData
	var validators []ValidatorKeyData
	if err := json.Unmarshal(output, &validators); err != nil {
		fmt.Printf("Error unmarshaling ValidatorKeyData: %v\n", err)
		os.Exit(1)
	}

	// Iterate over the validators to find the matching address
	found := false
	for _, validator := range validators {
		if validator.Address == address {
			// Decode the nested PubKey
			var pubKey PubKey
			if err := json.Unmarshal([]byte(validator.Pubkey), &pubKey); err != nil {
				fmt.Printf("Error unmarshaling PubKey: %v\n", err)
				os.Exit(1)
			}

			// Replace the Pubkey string with a decoded PubKey struct
			validator.Pubkey = pubKey.Key
			found = true
			// Return the matched ValidatorKeyData
			return validator
		}
	}

	// If no validator matches the address, exit with an error
	fmt.Printf("No validator found with address: %s\n", address)
	if !found {
		os.Exit(1)
	}

	return ValidatorKeyData{}
}
