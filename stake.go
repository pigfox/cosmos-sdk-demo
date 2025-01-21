package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func stake(validatorAddress string) {
	fmt.Println("Step 7: Delegating stake")
	fmt.Print("initial validatorAddress: ", validatorAddress, "\n")

	// Ensure the validator address has the correct 'cosmosvaloper' prefix
	if !strings.HasPrefix(validatorAddress, "cosmosvaloper") {
		fmt.Println("Error: Validator address is missing the correct 'cosmosvaloper' prefix.")
		os.Exit(1)
	}

	// Construct the command
	cmd := exec.Command(
		"simd", "tx", "staking", "delegate", validatorAddress, settings.Amount,
		"--from", settings.KeyName,
		"--chain-id", settings.ChainID,
		"--home", settings.AppHomeDir,
		"--keyring-backend", settings.KeyringBackend,
		"--broadcast-mode", "sync",
		"--yes",
		"--fees", settings.Fees,
	)

	// Capture the output of the command
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error delegating stake: %v\n", err)
		fmt.Println(out.String())
		os.Exit(1)
	}

	// Print the result
	fmt.Println("Delegation result:")
	fmt.Println(out.String())
}
