package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func node(state string) {
	if state == START {
		fmt.Println("Starting node...")
		// Start the node with the simd command
		cmd := exec.Command("simd", "start")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		// Execute the command
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error starting node: %v\n", err)
			fmt.Println(out.String())
			os.Exit(1)
		}
		fmt.Println("Node started successfully.")

	} else if state == STOP {
		fmt.Println("Stopping node...")
		// Stopping the node can be done by killing the process or using a specific command
		// Here we'll attempt to find the process ID (PID) of the running node and kill it
		cmd := exec.Command("pkill", "simd")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		// Execute the command
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error stopping node: %v\n", err)
			fmt.Println(out.String())
			os.Exit(1)
		}
		fmt.Println("Node stopped successfully.")

	} else {
		fmt.Println("Invalid state. Use 'start' or 'stop'.")
	}

	// Query the node's status
	cmd := exec.Command("simd", "status", "--node", "http://localhost:26657")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Execute the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error querying node status: %v\n", err)
		fmt.Println(out.String())
		os.Exit(1)
	}

	// Print the node's status
	fmt.Println("Node status:")
	fmt.Println(out.String())
}
