package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func getValidatorPubKey(accountAddress string) PubKey {
	fmt.Println("getValidatorPubKey()")

	validatorPubKeyCmd := []string{
		"keys", "list",
		"--home", settings.AppHomeDir,
		"--keyring-backend", settings.KeyringBackend,
		"--output", "json",
	}
	fmt.Println("Fetching keys with command:", validatorPubKeyCmd)

	cmd := exec.Command("simd", validatorPubKeyCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	var keys []ValidatorKeyData
	if err := json.Unmarshal(output, &keys); err != nil {
		fmt.Printf("Error unmarshaling ValidatorKeyData: %v\n", err)
		os.Exit(1)
	}

	for _, key := range keys {
		if key.Address == accountAddress {
			var pubKey PubKey
			if err := json.Unmarshal([]byte(key.Pubkey), &pubKey); err != nil {
				fmt.Printf("Error unmarshaling PubKey: %v\n", err)
				os.Exit(1)
			}
			return pubKey
		}
	}

	fmt.Printf("No public key found for address: %s\n", accountAddress)
	os.Exit(1)
	return PubKey{}
}
