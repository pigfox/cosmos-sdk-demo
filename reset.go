package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func reset() {
	fmt.Println("Reset the blockchain")

	// Remove old data directory
	dataDir := settings.AppHomeDir
	err := os.RemoveAll(dataDir)
	if err != nil {
		fmt.Printf("Failed to remove data directory: %v\n", err)
		os.Exit(1)
	}

	// Reinitialize the blockchain
	cmd := exec.Command("simd", "init", settings.Moniker, "--home", settings.AppHomeDir)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to reset blockchain: %v\nOutput: %s", err, out.String())
		os.Exit(1)
	}
	fmt.Println("Blockchain reset successfully!")

	// Start the blockchain
	cmd = exec.Command("simd", "start", "--home", settings.AppHomeDir)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start blockchain: %v\nOutput: %s", err, out.String())
		os.Exit(1)
	}
	fmt.Println("Blockchain started successfully!")
}
