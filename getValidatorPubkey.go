package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jimlawless/whereami"
)

func getValidatorPubkey() string {
	fmt.Println("Step 5: Get the validator public key")

	// Run the command to get the validator public key
	cmd := exec.Command("simd", "tendermint", "show-validator")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: Failed to retrieve validator public key.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	// Trim the output to clean up any extra newlines or spaces
	validatorPubkey := strings.TrimSpace(string(output))
	if validatorPubkey == "" {
		fmt.Println("Error: Validator public key is empty.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	// Print the validator public key
	fmt.Printf("Validator public key: %s\n", validatorPubkey)

	// Return the public key for further use (if necessary)
	return validatorPubkey
}
