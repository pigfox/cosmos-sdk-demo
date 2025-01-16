package main

import (
	"fmt"
)

const (
	KEYRING_BACKEND = "test"
	KEY_NAME        = "the-key"
	CHAIN_ID        = "1234567890"
	VALIDATOR_NAME  = "pigfox"
	HOME_DIR        = "./"
)

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
	step7(validatorAddress)
}
