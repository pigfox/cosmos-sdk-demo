package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AddValidatorAndKey adds a validator and the associated key based on the provided constants and stores the result in a Validator struct.
func addValidator() (string, string) {
	fmt.Println("addValidator()")

	// Step 1: Validate the key exists
	validateKeyCmd := []string{
		"keys", "list",
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--output", "json",
	}
	fmt.Println("Validating key existence with command:", validateKeyCmd)
	cmd := exec.Command("simd", validateKeyCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to list keys: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	var keys []map[string]interface{}
	if err := json.Unmarshal(output, &keys); err != nil {
		fmt.Printf("Failed to parse key list: %s\n", err)
		os.Exit(1)
	}

	keyExists := false
	for _, key := range keys {
		if key["name"] == settings.KeyName {
			keyExists = true
			break
		}
	}

	if !keyExists {
		fmt.Printf("Key '%s' not found in the keyring.\n", settings.KeyName)
		os.Exit(1)
	}

	// Step 2: Fetch the regular account address
	addValidatorCmd := []string{
		"keys", "show", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--address",
	}
	fmt.Println("Fetching regular account address with command:", addValidatorCmd)
	cmd = exec.Command("simd", addValidatorCmd...)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to fetch regular account address: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	accountAddress := strings.TrimSpace(string(output))
	fmt.Println("Regular account address fetched:", accountAddress)

	// Step 3: Fetch the validator address
	addValidatorCmd = []string{
		"keys", "show", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--bech", "val",
		"--address",
	}
	fmt.Println("Fetching validator address with command:", addValidatorCmd)
	cmd = exec.Command("simd", addValidatorCmd...)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to fetch validator address: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	validatorAddress := strings.TrimSpace(string(output))
	fmt.Println("Validator address fetched:", validatorAddress)

	return accountAddress, validatorAddress
}

func addValidatorFile(validatorAddress, pubKey string) {
	// Create the validator data structure
	validator := ValidatorFile{
		Name:        settings.Moniker,
		Address:     validatorAddress,
		PubKey:      pubKey,
		KeyringPath: settings.KeyringBackend,
		HomeDir:     settings.AppHomeDir,
	}

	// Serialize the validator data to JSON
	data, err := json.MarshalIndent(validator, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal validator data: %s", err)
		os.Exit(1)
	}

	// Write the JSON data to a file
	file, err := os.Create(settings.ValidatorPath)
	if err != nil {
		fmt.Printf("failed to create file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		fmt.Printf("failed to write to file: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Validator file created successfully: %s\n", settings.ValidatorPath)
}
