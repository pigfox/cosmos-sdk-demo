package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func getValidatorPubKey() ValidatorKeyData {
	fmt.Println("getValidatorPubKey()")
	validatorPubKeyCmd := []string{
		"keys", "list",
		"--home", settings.AppHomeDir,
		"--keyring-backend", settings.KeyringBackend,
		"--output", "json",
	}

	cmd := exec.Command("simd", validatorPubKeyCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	var validators []ValidatorKeyData
	if err := json.Unmarshal(output, &validators); err != nil {
		fmt.Printf("Error unmarshaling ValidatorKeyData: %v\n", err)
		os.Exit(1)
	}

	// Process each validator and decode the nested PubKey
	for _, validator := range validators {
		// Decode the nested PubKey
		var pubKey PubKey
		if err := json.Unmarshal([]byte(validator.Pubkey), &pubKey); err != nil {
			fmt.Printf("Error unmarshaling PubKey: %v\n", err)
			continue
		}

	}
	return validators[0]
}
