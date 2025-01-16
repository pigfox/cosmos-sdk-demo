package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func step0() {
	clear()
	fmt.Println("Step 0a: Reset the blockchain")

	// Define the command to reset the blockchain
	cmd := exec.Command("ignite", "chain", "reset")

	// Run the command and capture any output or errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to reset blockchain: %v\nOutput: %s", err, string(output))
		os.Exit(1)
	}

	fmt.Println("Blockchain reset successfully!")
	deleteKeys()
}

// clear clears the terminal screen
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command and handle errors
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to clear the terminal: %v\n", err)
	}
}

func deleteKeys() {
	fmt.Println("Step 0b: Checking keys")

	// Define the command to list keys in JSON format
	cmd := exec.Command("simd", "keys", "list", "--home", "./", "--keyring-backend", "test", "--output", "json")

	// Capture the output of the command
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("No keys to delete.")
		return
	}

	// Parse the output to extract key names
	output := out.String()
	keyNames := parseKeyNames(output)

	if len(keyNames) == 0 {
		fmt.Println("No keys to delete.")
		return
	}

	// Delete each key
	for _, key := range keyNames {
		fmt.Printf("Deleting key: %s\n", key)

		// Define the command to delete the key
		delCmd := exec.Command("simd", "keys", "delete", key, "--home", "./", "--keyring-backend", "test", "-y")

		// Capture the output of the delete command
		var delOut bytes.Buffer
		delCmd.Stdout = &delOut
		delCmd.Stderr = &delOut

		// Execute the delete command
		delErr := delCmd.Run()
		if delErr != nil {
			fmt.Printf("Failed to delete key %s: %v\n", key, delErr)
		}
	}

	// List remaining keys
	fmt.Println("Remaining keys:")
	listCmd := exec.Command("simd", "keys", "list", "--home", "./", "--keyring-backend", "test")
	listCmd.Stdout = &out
	listCmd.Stderr = &out
	_ = listCmd.Run()
	fmt.Println(out.String())
}

// parseKeyNames parses the JSON output from `simd keys list` and extracts key names
func parseKeyNames(jsonOutput string) []string {
	lines := strings.Split(jsonOutput, "\n")
	var keyNames []string
	for _, line := range lines {
		if strings.Contains(line, "name") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				// Extract the key name and trim quotes/spaces
				keyName := strings.Trim(parts[1], ` ",`)
				keyNames = append(keyNames, keyName)
			}
		}
	}
	return keyNames
}
