package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func stake(validatorAddress string) {
	fmt.Println("Step 7: Delegating stake")
	fmt.Print("initial validatorAddress: ", validatorAddress, "\n")

	// Ensure the validator address has the correct 'cosmosvaloper' prefix
	if !strings.HasPrefix(validatorAddress, "cosmosvaloper") {
		fmt.Println("Error: Validator address is missing the correct 'cosmosvaloper' prefix.")
		return
	}

	// Construct the command
	cmd := exec.Command(
		"simd", "tx", "staking", "delegate", validatorAddress, AMOUNT,
		"--from", keyName,
		"--chain-id", CHAIN_ID,
		"--home", APP_HOME_DIR,
		"--keyring-backend", KEYRING_BACKEND,
		"--broadcast-mode", "sync",
		"--yes",
		"--fees", FEES,
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
		return
	}

	// Print the result
	fmt.Println("Delegation result:")
	fmt.Println(out.String())
}
