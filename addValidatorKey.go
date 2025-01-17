package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jimlawless/whereami"
)

func addValidatorKey() {
	fmt.Println("addValidatorKey: Add a validator key if it doesn't exist")

	// Check if the validator key exists using test keyring backend
	cmd := exec.Command("simd", "keys", "show", VALIDATOR_NAME, "--keyring-backend", KEYRING_BACKEND)
	if err := cmd.Run(); err != nil {
		fmt.Println("Validator key not found. Adding validator key...", err, whereami.WhereAmI())

		// Add validator key without interactive input
		addCmd := exec.Command("simd", "keys", "add", KEY_NAME, "--keyring-backend", KEYRING_BACKEND)
		addCmd.Stdin = strings.NewReader("y\n") // Accept prompt automatically
		addCmd.Stdout = os.Stdout
		addCmd.Stderr = os.Stderr

		// Run the add command
		if err := addCmd.Run(); err != nil {
			fmt.Println("Error adding validator key: ", err, whereami.WhereAmI())
			os.Exit(1)
		}
		fmt.Println("Validator key added successfully!")
	} else {
		fmt.Println("Validator key exists.")
	}
}
