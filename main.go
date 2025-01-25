package main

import (
	"fmt"
)

var settings Settings

func main() {
	fmt.Println("Starting setup...")
	settings = newSettings()
	reset()
	simd()
	tools()
	initChain()
	acct1 := newAccount(ACCT1)
	acct1.AccountKey = addKey(acct1.KeyName)
	fmt.Printf("Account: %+v\n", acct1)
	acct2 := newAccount(ACCT2)
	acct2.AccountKey = addKey(acct2.KeyName)
	fmt.Printf("Account: %+v\n", acct2)
	validator := newAccount(VALIDATOR)
	validator.AccountKey = addKey(validator.KeyName)
	validator.AccountKey.Address = addValidatorKey(validator.KeyName)
	fmt.Printf("Account: %+v\n", validator)
	addGenesisFile(acct1, acct2, validator)
	/*



		addValidatorFile(validatorAddress, validatorPubKey.Key)

		node(START)
		stake(validatorAddress)
		node(STOP)
	*/
	fmt.Println("Run complete.")
}
