package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/cosmos/btcutil/bech32"
	"github.com/jimlawless/whereami"
)

func getValidatorAddress() string {
	fmt.Println("Step 4: Get the validator address")

	// Step 4.1: Get the account address
	cmd := exec.Command("simd", "keys", "show", VALIDATOR_NAME, "-a", "--keyring-backend", KEYRING_BACKEND)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error retrieving account address:", err, whereami.WhereAmI())
	}

	rawOutput := string(output)
	accountAddress := strings.TrimSpace(rawOutput)
	fmt.Printf("Raw account address output: %q\n", rawOutput)
	fmt.Printf("Trimmed account address: %q\n", accountAddress)

	if accountAddress == "" {
		log.Fatalf("Error: Account address is empty.")
	}

	// Step 4.2: Convert to validator address manually
	validatorAddress := convertToValidatorAddress(accountAddress)
	fmt.Printf("Validator address: %s\n", validatorAddress)

	return validatorAddress
}

func convertToValidatorAddress(accountAddress string) string {
	// Decode the account address using the Bech32 library
	hrp, data, err := bech32.Decode(accountAddress, 90)
	if err != nil {
		log.Fatalf("Error decoding account address:", accountAddress, err, whereami.WhereAmI())
	}

	fmt.Printf("Decoded HRP: %s, Decoded Data: %v\n", hrp, data)

	// Ensure the decoded data is not empty
	if len(data) == 0 {
		log.Fatalf("Error: Decoded data is empty for account address %s", accountAddress)
	}

	// Convert to validator address
	validatorPrefix := "cosmosvaloper" //"cosmos"
	validatorAddress, err := bech32.Encode(validatorPrefix, data)
	if err != nil {
		log.Fatalf("Error encoding validator address:", err, whereami.WhereAmI())
	}

	fmt.Printf("Encoded Validator Address: %s\n", validatorAddress)
	return validatorAddress
}
