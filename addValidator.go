package main

import (
	"fmt"
	"os"
	"os/exec"
)

// AddValidatorAndKey adds a validator and the associated key based on the provided constants and stores the result in a Validator struct.
func addValidator() string {
	fmt.Println("addValidator()")
	/*
		fmt.Println("addValidatorAndKey: Add a new validator and key")
		fmt.Println(settings)

		// Step 1: Add the key to the keyring
		addKeyCmd := []string{
			"keys", "add", settings.KeyName,
			"--keyring-backend", settings.KeyringBackend,
			"--home", settings.AppHomeDir,
			"--no-backup",
			"--log_level", "trace",
			"--output", "json",
		}

		fmt.Println("Adding key:", addKeyCmd)

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
	*/
	// Step 2: Create the validator using the provided parameters
	addValidatorCmd := []string{
		"keys", "show", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--bech", "val",
		"--address",
	}
	fmt.Println("addValidatorCmd:", addValidatorCmd)
	//$ (main) simd keys show my-key --keyring-backend test --home /home/peter/.simapp --bech val --address
	//cmd := exec.Command("simd", "tx", "staking", "create-validator", settings.ValidatorPath, "--from", settings.KeyName)
	cmd := exec.Command("simd", addValidatorCmd...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to create validator: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	validatorAddress := string(output)

	fmt.Println("Validator created successfully:", validatorAddress)

	/*
		$ (main) simd keys show my-key --keyring-backend test --home /home/peter/.simapp --address
		cosmos1rk2uueefpfzajrvjtaxerqclhz2aery4qa45cz

		$ (main) simd keys show my-key --keyring-backend test --home /home/peter/.simapp --bech val --address
		cosmosvaloper1rk2uueefpfzajrvjtaxerqclhz2aery49fpp53

		$ (main) simd query staking validator cosmos1rk2uueefpfzajrvjtaxerqclhz2aery4qa45cz
		post failed: Post "http://localhost:26657": dial tcp 127.0.0.1:26657: connect: connection refused
	*/
	return validatorAddress
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
