package main

import (
	"fmt"
	"os"
	"strconv"
)

// Transfer funds from acct1 to acct2
func transfer(acct1, acct2 Account, amount int) {
	// Step 1: Build the transfer command
	transferCmd := []string{
		"tx", "send", acct1.Details.Address, acct2.Details.Address,
		strconv.Itoa(amount),
		"--keyring-backend", settings.KeyringBackend, // You can adjust this to your keyring backend
		"--home", settings.AppHomeDir, // Path to your application home
		"--yes",            // Auto-confirmation
		"--output", "json", // Output format
	}

	// Step 2: Execute the transfer command using simdCmd
	output, err := simdCmd(transferCmd)
	if err != nil {
		fmt.Printf("Failed to send transaction: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}

	// Step 3: Print the output (transaction result)
	fmt.Println("Transfer successful!")
	fmt.Println(string(output))

}
