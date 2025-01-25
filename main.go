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
	initChain()
	regularKey := addRegularKey()
	/*
		regularKey2 := addRegularKey("2")
		if regularKey == regularKey2 {
			fmt.Println("Error: Same keys.")
			os.Exit(1)
		}
	*/
	fmt.Println("Regular key:", regularKey)
	accountAddress := getAccountAddress()
	validatorAddress := addValidator()
	fmt.Println("Regular account address:", accountAddress)
	fmt.Println("Validator address:", validatorAddress)

	validatorPubKey := getValidatorPubKey(accountAddress)
	fmt.Println("Validator public key:", validatorPubKey)

	addGenesisFile(accountAddress, validatorAddress, validatorPubKey)

	addValidatorFile(validatorAddress, validatorPubKey.Key)

	node(START)
	stake(validatorAddress)
	node(STOP)
	/**/
	fmt.Println("Run complete.")
}
