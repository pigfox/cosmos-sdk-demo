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
	fmt.Println("Regular key:", regularKey)
	accountAddress := getAccountAddress()
	//my_validator_address=$(simd keys show "$KEY_NAME" --keyring-backend "$KEYRING_BACKEND" --home "$HOME_DIR" --bech val --address)
	validatorAddress := addValidator()
	fmt.Println("Regular account address:", accountAddress)
	fmt.Println("Validator address:", validatorAddress)

	validatorPubKey := getValidatorPubKey(accountAddress)
	fmt.Println("Validator public key:", validatorPubKey)

	addGenesisFile(accountAddress, validatorAddress, validatorPubKey)

	addValidatorFile(validatorAddress, validatorPubKey.Key)
	/*
		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
	fmt.Println("Run complete.")
}
