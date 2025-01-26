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

	fmt.Println(output)

	// Step 2: Define a temporary struct to handle the raw pubkey field
	type TempAccount struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Address string `json:"address"`
		PubKey  string `json:"pubkey"` // PubKey as a raw string
	}

	var tempAccount TempAccount
	if err := json.Unmarshal([]byte(output), &tempAccount); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Step 3: Parse the raw pubkey string into the PubKey struct
	var pubKey PubKey
	if err := json.Unmarshal([]byte(tempAccount.PubKey), &pubKey); err != nil {
		log.Fatalf("Error unmarshaling PubKey: %v", err)
	}

	// Step 4: Assemble the AccountKey with the parsed PubKey
	account := AccountKey{
		Name:    tempAccount.Name,
		Type:    tempAccount.Type,
		Address: tempAccount.Address,
		Public:  pubKey,
	}

	return account
}
