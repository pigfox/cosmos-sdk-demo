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

	fmt.Println("Data directory removed:", dir)

	cmd := exec.Command("mkdir", "-p", dir)
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to create data directory: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Data directory created:", dir)
}

func clear() {
	fmt.Println("Clearing the screen...")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
