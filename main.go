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
	addRegularKey()
	accountAddress, validatorAddress := addValidator()
	fmt.Println("Regular account address:", accountAddress)
	fmt.Println("Validator address:", validatorAddress)
	//validatorAddress := addValidator()
	fmt.Println("Validator address:", validatorAddress)
	validatorPubKey := getValidatorPubKey(accountAddress)
	fmt.Println("Validator public key:", validatorPubKey)
	addGenesis(accountAddress, validatorAddress, validatorPubKey)
	/*
		node(START)
		stake(validatorAddress)
		node(STOP)

	*/
}
