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
	settings = newSettings()
	clearSetup()
	simd()
	tools()
	addRegularAccount()
	addValidatorAndKey() //<---bug
	/*
		os.Exit(0)
		addGenesis(accountAddress, validatorAddress, validatorPubkey)
		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
}
