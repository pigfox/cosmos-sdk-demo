package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func addRegularKey() {
	fmt.Println("addRegularKey()")
	fmt.Println("addValidatorAndKey: Add a new validator and key")
	fmt.Println(settings)

	// Step 1: Add the key to the keyring
	addKeyCmd := []string{
		"keys", "add", settings.KeyName,
		"--keyring-backend", settings.KeyringBackend,
		"--home", settings.AppHomeDir,
		"--no-backup",
		"--log_level", "trace",
		"--output", "json",
	}

	fmt.Println("Adding key:", addKeyCmd)

	cmd := exec.Command("simd", addKeyCmd...)
	cmd.Stdin = bytes.NewReader([]byte("y\n"))

	// Capture combined output (stdout + stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to add key: %s\n", err)
		fmt.Printf("Command Output: %s\n", string(output))
		os.Exit(1)
	}
	fmt.Println("Key added successfully:", string(output))
}
