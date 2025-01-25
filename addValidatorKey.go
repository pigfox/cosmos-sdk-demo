package main

import (
	"fmt"
	"os"
)

func addValidatorKey(keyName string) string {
	fmt.Println("validatorKey()")

	// Step 1: Add the key to the keyring
	showKeyCmd := []string{
		"keys", "show", keyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--bech", "val",
		"--address",
	}

	output, err := simdCmd(showKeyCmd)
	if err != nil {
		fmt.Printf("Failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", output)
		os.Exit(1)
	}

	return output
}
