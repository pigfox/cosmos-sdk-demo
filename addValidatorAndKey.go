package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AddValidatorAndKey adds a validator and the associated key based on the provided constants and stores the result in a Validator struct.
func addValidatorAndKey() {
	fmt.Println("addValidatorAndKey: Add a new validator and key")

	// Step 1: Add the key to the keyring
	simdPath := getHomeDir() + "/.simapp" // Replace with actual path to simd binary
	addKeyCmd := fmt.Sprintf(
		"simd keys add %s --keyring-backend %s --key-type secp256k1 --no-backup --home %s --account %d --interactive=false --coin-type 118",
		settings.KeyName, settings.KeyringBackend, simdPath, 0,
	)

	fmt.Println("Adding key:", addKeyCmd)
	cmd := exec.Command(addKeyCmd)
	cmd.Dir = settings.AppHomeDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	fmt.Println("Key added successfully:", string(output))

	// Step 2: Create the validator using the provided parameters
	createValidatorCmd := fmt.Sprintf(
		"simd tx staking create-validator --amount=%s --pubkey=%s --moniker=%s --chain-id=%s --from=%s --fees=%s --gas=auto --yes",
		settings.Amount, "cosmos-sdk", settings.ValidatorName, settings.ChainID, settings.KeyName, settings.Fees,
	)

	fmt.Println("Creating validator:", createValidatorCmd)
	cmd = exec.Command("bash", "-c", createValidatorCmd)
	cmd.Dir = settings.AppHomeDir
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to create validator: %s\n", err)
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
