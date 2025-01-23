package main

import (
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

	initCmd := []string{
		"keys", "init",
		settings.Moniker,
		"--home", settings.AppHomeDir,
	}
	out, err := simdCmd(initCmd)
	if err != nil {
		fmt.Printf("Failed to reset blockchain: %v\nOutput: %s", err, out)
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
