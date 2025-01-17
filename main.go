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
)

var keyName = ""

func main() {
	fmt.Println("Starting setup...")

	//source := "./genesis.json"
	//destination := getHomeDir() + "/.simapp/config/genesis.json"
	step0()
	step1()
	step2()
	step3()
	validatorAddress := step4()
	validatorPubkey := step5()
	step6(validatorAddress, validatorPubkey)
	node("start")
	step7(validatorAddress)
	node("stop")
}
