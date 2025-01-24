package main

import (
	"fmt"
	"os"
	"os/exec"
)

func initChain() {
	fmt.Println("initChain()")
	// Step 1: Initialize the blockchain
	initChainCmd := []string{
		"init",
		settings.Moniker,
		"--chain-id", settings.ChainID,
		"--home", settings.AppHomeDir,
	}

	cmd := exec.Command("simd", initChainCmd...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to initialize the blockchain: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Blockchain initialized successfully.")
}
