package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func addKey(keyName string) AccountKey {
	fmt.Println("addKey()")

	// Step 1: Add the key to the keyring
	addKeyCmd := []string{
		"keys", "add", keyName,
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

	var account AccountKey
	err = json.Unmarshal([]byte(output), &account)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return account
}
