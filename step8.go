package main

import (
	"fmt"
	"os"
)

// Helper function to copy a file
func step8(src, dst string) {
	fmt.Println("Step 8: Copy the genesis file to the target directory")

	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error: Failed to open the source file:", err)
		os.Exit(1)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error: Failed to create the destination file:", err)
		os.Exit(1)
	}
	defer destinationFile.Close()

	_, err = sourceFile.Stat()
	if err != nil {
		fmt.Println("Error: Failed to get the source file info:", err)
		os.Exit(1)
	}

	_, err = destinationFile.Stat()
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error: Failed to get the destination file info:", err)
		os.Exit(1)
	}

	_, err = destinationFile.ReadFrom(sourceFile)
	if err != nil {
		fmt.Println("Error: Failed to copy the source file to the destination file:", err)
		os.Exit(1)
	}

	fmt.Printf("Genesis file copied to %s successfully.\n", dst)
}
