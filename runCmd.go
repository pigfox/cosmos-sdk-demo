package main

import "os/exec"

func runCommand(cmdArgs []string) (string, error) {
	cmd := exec.Command("simd", cmdArgs...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
