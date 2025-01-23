package main

import (
	"fmt"
)

const (
	START = "start"
	STOP  = "stop"
)

var settings Settings

func main() {
	fmt.Println("Starting setup...")
	settings = newSettings()
	reset()
	simd()
	tools()
	addRegularKey()
	accountAddress, validatorAddress := addValidator()
	fmt.Println("Regular account address:", accountAddress)
	fmt.Println("Validator address:", validatorAddress)
	validatorPubKey := getValidatorPubKey(accountAddress)
	fmt.Println("Validator public key:", validatorPubKey)
	addGenesis(accountAddress, validatorAddress, validatorPubKey)
	createValidatorFile(validatorAddress, validatorPubKey.Key)
	node(START)
	stake(validatorAddress)
	node(STOP)
	fmt.Println("Run complete.")
}
