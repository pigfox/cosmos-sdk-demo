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
	acct1 := newAccount("acct1")
	acct1.AccountKey = addKey(acct1.KeyName)
	fmt.Printf("Account: %+v\n", acct1)
	acct2 := newAccount("acct2")
	acct2.AccountKey = addKey(acct2.KeyName)
	fmt.Printf("Account: %+v\n", acct2)
	validator := newAccount("validator")
	validator.AccountKey = addKey(validator.KeyName)
	validator.AccountKey.Address = addValidatorKey(validator.KeyName)
	fmt.Println(validator)
	addGenesisFile(acct1, acct2, validator)
	/*



		addValidatorFile(validatorAddress, validatorPubKey.Key)

		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
	fmt.Println("Run complete.")
}
