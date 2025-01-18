package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func clearSetup() {
	clear()
	fmt.Println("clearSetup: Reset the blockchain")

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
	fmt.Println("Step 0b: Deleting all keys")

	// Remove all files in the keyring-test directory
	err := os.RemoveAll("keyring-test/*")
	if err != nil {
		fmt.Printf("Error clearing keyring-test directory: %v\n", err)
		return
	}

	// Define the command to list keys in JSON format
	cmd := exec.Command("simd", "keys", "list", "--home", APP_HOME_DIR, "--keyring-backend", KEYRING_BACKEND, "--output", "json")

	// Capture the output of the command
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Execute the command
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error listing keys:", err)
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
		delCmd := exec.Command("simd", "keys", "delete", key, "--home", APP_HOME_DIR, "--keyring-backend", KEYRING_BACKEND, "-y")

		// Capture the output of the delete command
		var delOut bytes.Buffer
		delCmd.Stdout = &delOut
		delCmd.Stderr = &delOut

		// Execute the delete command
		delErr := delCmd.Run()
		if delErr != nil {
			fmt.Printf("Failed to delete key %s: %v\n", key, delErr)
		} else {
			fmt.Printf("Key %s deleted successfully.\n", key)
		}
	}

	time.Sleep(2 * time.Second)

	// List remaining keys
	fmt.Println("Remaining keys:")
	listCmd := exec.Command("simd", "keys", "list", "--home", APP_HOME_DIR, "--keyring-backend", KEYRING_BACKEND)
	listCmd.Stdout = &out
	listCmd.Stderr = &out
	_ = listCmd.Run()
	fmt.Println(out.String())
}

// parseKeyNames parses the JSON output from `simd keys list` and extracts key names
func parseKeyNames(jsonOutput string) []string {
	// Regular expression to match key names in the output
	re := regexp.MustCompile(`"name":\s*"([^"]+)"`)
	matches := re.FindAllStringSubmatch(jsonOutput, -1)

	var keyNames []string
	for _, match := range matches {
		// match[1] contains the captured key name
		keyNames = append(keyNames, match[1])
	}

	return keyNames
}
