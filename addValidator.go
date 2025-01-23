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

func createValidatorFile(validatorKeyData ValidatorKeyData, pubKey PubKey) {
	// Create a validator file
	validatorJSON := `{
		"address": "` + validatorKeyData.Address + `",
		"pub_key": {
			"type": "` + pubKey.Type + `",
			"value": "` + pubKey.Key + `"}`

	err := os.WriteFile(settings.ValidatorPath, []byte(validatorJSON), 0644)
	if err != nil {
		fmt.Printf("Failed to write validator JSON file: %s\n", err)
		os.Exit(1)
	}
}

/*
	validatorData := map[string]interface{}{
		"name":     settings.ValidatorName,
		"pubkey":   settings.ValidatorPubKey,
		"amount":   settings.StakeAmount,
		"moniker":  settings.ValidatorMoniker,
		"chain-id": settings.ChainID,
		"fees":     settings.Fees,
		"gas":      "auto",
	}
*/
