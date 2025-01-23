package main

import "os/exec"

func simdCmd(cmdArgs []string) (string, error) {
	cmd := exec.Command("simd", cmdArgs...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
