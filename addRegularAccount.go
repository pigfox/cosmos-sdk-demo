package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func addRegularAccount() {
	fmt.Println("addAccount: Add a new account")
	fmt.Println("KEY_NAME", KEY_NAME)
	fmt.Println("KEYRING_BACKEND", KEYRING_BACKEND)

	// Execute the simd command
	cmd := exec.Command("simd", "keys", "add", KEY_NAME, "--home", APP_HOME_DIR, "--keyring-backend", KEYRING_BACKEND)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	// Print the output from the command
	//fmt.Println("Key details:")
	//fmt.Println(string(output))
	parseAccountDetails(string(output))
}

func parseAccountDetails(output string) {
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "address:") {
			regularAccount.Address = strings.TrimSpace(strings.TrimPrefix(line, "address:"))
		} else if strings.HasPrefix(line, "name:") {
			regularAccount.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
		} else if strings.HasPrefix(line, "pubkey:") {
			regularAccount.PubKey = strings.TrimSpace(strings.TrimPrefix(line, "pubkey:"))
		} else if strings.HasPrefix(line, "type:") {
			regularAccount.Type = strings.TrimSpace(strings.TrimPrefix(line, "type:"))
		}
	}

	if regularAccount.Address == "" || regularAccount.Name == "" || regularAccount.PubKey == "" || regularAccount.Type == "" {
		fmt.Println("failed to parse account details: incomplete data")
		fmt.Printf("+%v", regularAccount)
		os.Exit(1)
	}
}
