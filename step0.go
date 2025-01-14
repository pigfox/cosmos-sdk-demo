package main

import (
	"fmt"
	"os"
	"os/exec"
)

func step0() {
	fmt.Println("Step 0: Reset the blockchain")

	// Define the command to reset the blockchain
	cmd := exec.Command("ignite", "chain", "reset")

	// Run the command and capture any output or errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to reset blockchain: %v\nOutput: %s", err, string(output))
		os.Exit(1)
	}

	fmt.Println("Blockchain reset successfully!")
}
