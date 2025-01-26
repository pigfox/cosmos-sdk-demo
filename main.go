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
	acct1.Details = addKey(acct1.Name)
	fmt.Printf("Account: %+v\n", acct1)
	acct2 := newAccount(ACCT2)
	acct2.Details = addKey(acct2.Name)
	fmt.Printf("Account: %+v\n", acct2)
	validator := newAccount(VALIDATOR)
	validator.Details = addKey(validator.Name)
	validator.Details.Address = addValidatorKey(validator.Name)
	fmt.Printf("Account: %+v\n", validator)
	addGenesisFile(acct1, acct2, validator)
	node(START)
	transfer(acct1, acct2, 50000)
	stake(validator.Details.Address)
	node(STOP)
	fmt.Println("Run complete.")
}
