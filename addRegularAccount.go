package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
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
	fmt.Println(string(output))
	parseAccountDetails(string(output))
}

func parseAccountDetails(output string) {
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip lines that are empty
		if len(line) == 0 {
			continue
		}

		// Remove the first two characters (dash and space) if they exist
		if len(line) > 1 && line[0] == '-' && line[1] == ' ' {
			line = line[2:]
		}

		// Debugging log to inspect the line
		fmt.Printf("Processing line: '%s'\n", line)

		if strings.HasPrefix(line, "address:") {
			address := strings.TrimSpace(strings.TrimPrefix(line, "address:"))
			// Debugging log
			fmt.Printf("Parsed address: '%s'\n", address)
			regularAccount.Address = address
		} else if strings.HasPrefix(line, "name:") {
			regularAccount.Name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
		} else if strings.HasPrefix(line, "pubkey:") {
			pubkeyStr := strings.TrimSpace(strings.TrimPrefix(line, "pubkey:"))
			pubkeyStr = strings.Trim(pubkeyStr, "'") // Remove surrounding single quotes if present

			// Parse the pubkey JSON string into the RegularAccountPubKey struct
			if err := json.Unmarshal([]byte(pubkeyStr), &regularAccount.RegularAccountPubKey); err != nil {
				fmt.Println("Error parsing pubkey:", err)
				os.Exit(1)
			}
		} else if strings.HasPrefix(line, "type:") {
			regularAccount.Type = strings.TrimSpace(strings.TrimPrefix(line, "type:"))
		}
	}

	// Check if any key details are missing
	if regularAccount.Address == "" || regularAccount.Name == "" || regularAccount.RegularAccountPubKey.Key == "" || regularAccount.Type == "" {
		fmt.Println("Failed to parse account details: incomplete data")
		printStructFields(regularAccount)
		os.Exit(1)
	}

	// Output the parsed regular account
	fmt.Printf("Parsed Account: %+v\n", regularAccount)
}

func printStructFields(s interface{}) {
	v := reflect.ValueOf(s)

	// If it's a pointer, get the value it points to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure we're working with a struct
	if v.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}

	// Iterate over the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		// Print the field name and value
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
