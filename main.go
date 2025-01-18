package main

import (
	"fmt"
)

const (
	KEYRING_BACKEND = "test"
	KEY_NAME        = "my-key"
	CHAIN_ID        = "my-chain"
	VALIDATOR_NAME  = "pigfox"
	FEES            = "10000stake"
	AMOUNT          = "500000stake"
	GENESIS_PATH    = "/.simapp/config/genesis.json"
	APP_HOME_DIR    = "./"
	START           = "start"
	STOP            = "stop"
)

var regularAccount RegularAccount
var validator Validator

func main() {
	fmt.Println("Starting setup...")
	clearSetup()
	simd()
	tools()
	addRegularAccount()
	addValidatorAndKey()
	/*
		os.Exit(0)
		addGenesis(accountAddress, validatorAddress, validatorPubkey)
		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
}
