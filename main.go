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

func newSettings() Settings {
	homeDir := getHomeDir()
	fmt.Println("Home directory:", homeDir)
	return Settings{
		KeyringBackend: "test",
		KeyName:        "my-key",
		ChainID:        "my-chain",
		ValidatorName:  "pigfox",
		Fees:           "10000stake",
		Amount:         "500000stake",
		GenesisPath:    homeDir + "/.simapp/config/genesis.json",
		AppHomeDir:     homeDir,
	}
}
