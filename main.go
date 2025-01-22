package main

import (
	"fmt"
)

const (
	START = "start"
	STOP  = "stop"
)

var regularAccount RegularAccount
var validatorKeyData ValidatorKeyData
var settings Settings

func main() {
	fmt.Println("Starting setup...")
	settings = newSettings()
	clearSetup()
	simd()
	tools()
	accountAddress := addRegularAccount()
	fmt.Println("Regular account address:", accountAddress)
	addRegularKey()
	validatorAddress := addValidator()
	fmt.Println("Validator address:", validatorAddress)

	/*
		addGenesis(accountAddress, validatorAddress, validatorPubkey)
			os.Exit(0)

			node(START)
			stake(validatorAddress)
			node(STOP)
	*/
}
