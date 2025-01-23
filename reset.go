package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func reset() {
	clear()
	fmt.Println("Reset the blockchain")

	dir := settings.AppHomeDir
	err := os.RemoveAll(dir)
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
}

func clear() {
	fmt.Println("Clearing the screen...")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
