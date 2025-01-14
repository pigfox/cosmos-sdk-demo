package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func step4(validatorName string) string {
	fmt.Println("Step 4: Get the validator address")

	// Run command to get validator address using test keyring backend
	cmd := exec.Command("simd", "keys", "show", validatorName, "-a", "--keyring-backend", "test")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error: Failed to retrieve validator address: %v", err)
	}

	// Trim and clean up output
	validatorAddress := strings.TrimSpace(string(output))
	if validatorAddress == "" {
		log.Fatalf("Error: Validator address is empty.")
	}

	fmt.Printf("Validator address: %s\n", validatorAddress)
	return strings.ToLower(validatorAddress)
}
