package main

import (
	"fmt"
)

const (
	KEYRING_BACKEND = "test"
	CHAIN_ID        = "my-chain"
	VALIDATOR_NAME  = "pigfox"
	FEES            = "10000stake"
	AMOUNT          = "500000stake"
	GENESIS_PATH    = "/.simapp/config/genesis.json"
	APP_HOME_DIR    = "./"
	START           = "start"
	STOP            = "stop"
)

var keyName = ""

func main() {
	fmt.Println("Starting setup...")
	clearSetup()
	simd()
	tools()
	addValidatorKey()
	validatorAddress := getValidatorAddress()
	validatorPubkey := getValidatorPubkey()
	accountAddress := getAccountAddress()
	addGenesis(accountAddress, validatorPubkey)
	node(START)
	stake(validatorAddress)
	node(STOP)
}
