package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AddValidatorAndKey adds a validator and the associated key based on the provided constants and stores the result in a Validator struct.
func addValidatorAndKey() {
	fmt.Println("addValidatorAndKey: Add a new validator and key")
	fmt.Println(settings)
	/*
		simd keys add my-key --keyring-backend file --home /home/peter/.simapp
		Enter keyring passphrase (attempt 1/3):
		password must be at least 8 characters
		Enter keyring passphrase (attempt 2/3):
		Re-enter keyring passphrase:
	*/
	// Step 1: Add the key to the keyring
	addKeyCmd := []string{
		"keys", "add", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--no-backup",
		"--log_level", "trace",
	}

	cmd := exec.Command("simd", addKeyCmd...)
	cmd.Stdin = bytes.NewReader([]byte("y\n"))

	// Capture combined output (stdout + stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	fmt.Println("Key added successfully:", string(output))

	// Step 2: Create the validator using the provided parameters
	/*
		Examples: $ simd tx staking create-validator path/to/validator.json --from keyname
	*/

	fmt.Println("Creating validator:", settings.ValidatorPath)
	cmd = exec.Command("simd", "tx", "staking", "create-validator", settings.ValidatorPath, "--from", settings.KeyName)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create validator: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	fmt.Println("Validator created successfully:", string(output))

	// Step 3: Extract the address and pubkey from the output
	addressIndex := strings.Index(string(output), `"address":`)
	pubKeyIndex := strings.Index(string(output), `"pubkey":`)
	if addressIndex == -1 || pubKeyIndex == -1 {
		fmt.Println("failed to extract address or pubkey from output")
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	// Extracting address and pubkey
	addressStart := strings.Index(string(output)[addressIndex:], `"address":`) + len(`"address":`) + 1
	addressEnd := strings.Index(string(output)[addressStart:], `"`) + addressStart
	validator.Address = string(output)[addressStart:addressEnd]

	pubKeyStart := strings.Index(string(output)[pubKeyIndex:], `"pubkey":`) + len(`"pubkey":`) + 1
	pubKeyEnd := strings.Index(string(output)[pubKeyStart:], `"`) + pubKeyStart
	validator.PubKey = string(output)[pubKeyStart:pubKeyEnd]

	if validator.Address == "" || validator.PubKey == "" {
		fmt.Println("failed to extract address or pubkey from output")
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	fmt.Println("Validator Details:", validator)
}
