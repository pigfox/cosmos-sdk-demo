package main

import (
	"fmt"
	"os"
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

	_, err := simdCmd(initChainCmd)
	if err != nil {
		fmt.Printf("Failed to initialize the blockchain: %s\n", err)
		os.Exit(1)
	}
}
