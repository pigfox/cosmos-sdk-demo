package main

import (
	"fmt"
)

const (
	START = "start"
	STOP  = "stop"
)

var regularAccount RegularAccount
var validator Validator
var settings Settings

func main() {
	fmt.Println("Starting setup...")
	clearSetup()
	settings = newSettings()
	fmt.Println("Settings:", settings)
	simd()
	tools()
	//addRegularAccount()
	//addValidatorAndKey()
	/*
		os.Exit(0)
		addGenesis(accountAddress, validatorAddress, validatorPubkey)
		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
}
