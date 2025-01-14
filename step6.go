package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func step6(validatorAddress string) string {
	fmt.Println("Step 6: Convert validator address to operator address")

	// Convert validator address to operator address
	cmd := exec.Command("simd", "keys", "parse", validatorAddress)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error: Failed to parse validator address.")
		os.Exit(1)
	}

	// Print the raw output for debugging
	//fmt.Println("Raw output from 'simd keys parse':", string(output))

	// Define a regular expression to extract the operator address
	re := regexp.MustCompile(`human:\s*(\S+)`)
	matches := re.FindStringSubmatch(string(output))

	if len(matches) < 2 {
		fmt.Println("Error: Failed to extract operator address.")
		os.Exit(1)
	}

	// The operator address is the first captured match
	validatorOperatorAddress := matches[1]
	fmt.Println("Validator operator address:", validatorOperatorAddress)

	return validatorOperatorAddress
}
