package main

import (
	"fmt"
	"os"
)

// Get the current user's home directory
func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Failed to retrieve the home directory:", err)
		os.Exit(1)
	}

	return homeDir
}
