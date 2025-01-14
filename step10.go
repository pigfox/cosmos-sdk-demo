package main

import (
	"fmt"
	"os"
	"os/exec"
)

func step10() {
	fmt.Println("Step 10: Validate the genesis file")

	// Command to validate the genesis file
	validateCmd := exec.Command("simd", "genesis", "validate", "--trace")
	validateCmd.Stdout = os.Stdout
	validateCmd.Stderr = os.Stderr

	// Run the command and check for errors
	err := validateCmd.Run()
	if err != nil {
		fmt.Println("Error: Genesis validation failed.", err)
		os.Exit(1)
	}

	fmt.Println("Genesis file validated successfully.")
}
