package main

import (
	"fmt"
	"os"
	"time"
)

func step9(destination string) {
	fmt.Println("Step 9: Check the genesis file time stamp in the target directory")

	// Get the file information
	fileInfo, err := os.Stat(destination)
	if err != nil {
		fmt.Printf("Error getting file information for %s: %v\n", destination, err)
		return
	}

	// Get the current time and file modification time
	currentTime := time.Now()
	modTime := fileInfo.ModTime()

	// Calculate the file's age in seconds
	ageInSeconds := int(currentTime.Sub(modTime).Seconds())

	// Print the details
	fmt.Printf("File: %s\n", destination)
	fmt.Printf("Last modified: %v\n", modTime)
	fmt.Printf("File age: %d seconds\n", ageInSeconds)
}
