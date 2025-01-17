package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jimlawless/whereami"
)

func simd() {
	fmt.Println("Step 1: Checking 'simd' requirements...")

	// Check if 'simd' command is available
	_, err := exec.LookPath("simd")
	if err != nil {
		fmt.Println("Error: 'simd' command not found. Please install the Cosmos SDK.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	// Print 'simd' version
	cmd := exec.Command("simd", "version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: Failed to get 'simd' version.", whereami.WhereAmI(), err)
		os.Exit(1)
	}
}
