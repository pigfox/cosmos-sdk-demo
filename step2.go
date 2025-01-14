package main

import (
	"fmt"
	"os"
	"os/exec"
)

func step2() {
	fmt.Println("Step 2a: Checking 'jq' requirements...")
	if _, err := exec.LookPath("jq"); err != nil {
		fmt.Println("Error: 'jq' command not found. Please install 'jq' for JSON processing.")
		os.Exit(1)
	}

	fmt.Println("Step 2b: Checking 'libsecret-tools' requirements...")
	if _, err := exec.LookPath("secret-tool"); err != nil {
		fmt.Println("Error: 'secret-tool' command not found. Please install 'libsecret-tools'.")
		os.Exit(1)
	}

	fmt.Println("Step 2c: Checking 'expect' requirements...")
	if _, err := exec.LookPath("expect"); err != nil {
		fmt.Println("Error: 'expect' command not found. Please install 'expect' for handling interactive shell scripts.")
		os.Exit(1)
	}

	fmt.Println("All required tools are installed.")
}
