package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func step7(validatorAddress string) {
	fmt.Println("Step 7: Delegating stake")

	// Define the parameters
	fees := "10000stake"
	amount := "500000stake"

	// Construct the command
	cmd := exec.Command(
		"simd", "tx", "staking", "delegate", validatorAddress, amount,
		"--from", KEY_NAME,
		"--chain-id", CHAIN_ID,
		"--home", HOME_DIR,
		"--keyring-backend", KEYRING_BACKEND,
		"--broadcast-mode", "sync",
		"--yes",
		"--fees", fees,
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
