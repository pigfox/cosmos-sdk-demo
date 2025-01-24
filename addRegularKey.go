package main

import (
	"fmt"
	"os"
)

func addRegularKey() string {
	fmt.Println("addRegularKey()")

	// Step 1: Add the key to the keyring
	addKeyCmd := []string{
		"keys", "add", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--no-backup",
		"--output", "json",
	}

	output, err := simdCmd(addKeyCmd)
	if err != nil {
		fmt.Printf("Failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	fmt.Println("Key added successfully:", string(output))

	//simd keys show "$KEY_NAME" --keyring-backend "$KEYRING_BACKEND" --home "$HOME_DIR" --address
	showKeyCmd := []string{
		"keys", "show", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--address",
	}

	output, err = simdCmd(showKeyCmd)
	if err != nil {
		fmt.Printf("Failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	return string(output)

}
