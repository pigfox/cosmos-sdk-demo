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

	vCmd := []string{"version"}
	out, err := simdCmd(vCmd)
	if err != nil {
		fmt.Println("Error: Failed to get 'simd' version.", err, out)
		os.Exit(1)
	}
	fmt.Println("simd version:", out)
}
