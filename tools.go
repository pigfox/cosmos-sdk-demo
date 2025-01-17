package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jimlawless/whereami"
)

func tools() {
	fmt.Println("tools a: Checking 'jq' requirements...")
	if _, err := exec.LookPath("jq"); err != nil {
		fmt.Println("Error: 'jq' command not found. Please install 'jq' for JSON processing.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	fmt.Println("tools b: Checking 'libsecret-tools' requirements...")
	if _, err := exec.LookPath("secret-tool"); err != nil {
		fmt.Println("Error: 'secret-tool' command not found. Please install 'libsecret-tools'.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	fmt.Println("tools c: Checking 'expect' requirements...")
	if _, err := exec.LookPath("expect"); err != nil {
		fmt.Println("Error: 'expect' command not found. Please install 'expect' for handling interactive shell scripts.", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	fmt.Println("All required tools are installed.")
}
