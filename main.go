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
	simd()
	tools()
	reset()
	addRegularKey()
	accountAddress, validatorAddress := addValidator()
	//fmt.Println("Regular account address:", accountAddress)
	//fmt.Println("Validator address:", validatorAddress)
	validatorPubKey := getValidatorPubKey(accountAddress)
	//fmt.Println("Validator public key:", validatorPubKey)
	addGenesisFile(accountAddress, validatorAddress, validatorPubKey)
	addValidatorFile(validatorAddress, validatorPubKey.Key)
	node(START)
	stake(validatorAddress)
	node(STOP)
	fmt.Println("Run complete.")
}
