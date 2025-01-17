package main

import (
	"fmt"
)

const (
	KEYRING_BACKEND = "test"
	CHAIN_ID        = "my-chain"
	VALIDATOR_NAME  = "pigfox"
	HOME_DIR        = "./"
	FEES            = "10000stake"
	AMOUNT          = "500000stake"
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
