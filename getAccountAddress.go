package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getAccountAddress() string {
	fmt.Println("getAccountAddress: Getting account address")
	// Prepare the simd command to fetch the address
	cmd := exec.Command("simd", "keys", "show", settings.KeyName, "-a", "--home", settings.AppHomeDir, "--keyring-backend", settings.KeyringBackend)

	// Execute the command and capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		msg := fmt.Errorf("failed to execute simd command: %v\nOutput: %s", err, out.String())
		fmt.Println(msg, err)
		os.Exit(1)
	}

	// Return the address from the output
	return strings.TrimSpace(out.String())
}
