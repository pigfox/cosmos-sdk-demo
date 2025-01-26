package main

import (
	"fmt"
	"github.com/jimlawless/whereami"
	"os"
	"regexp"
)

func accountValidation(gp GenesisParams) {
	if gp.Validator.Name == VALIDATOR {
		regex := `^cosmosvaloper1[a-z0-9]{38}$`
		matched, err := regexp.MatchString(regex, gp.Validator.Details.Address)
		if err != nil {
			fmt.Println("Error with regex:", whereami.WhereAmI(), err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("Error: validator address is not in the correct format", whereami.WhereAmI())
			fmt.Printf("validatorAddress (trimmed): %q, length: %d\n", gp.Validator.Details.Address, len(gp.Validator.Details.Address))
			os.Exit(1)
		}
	}

	if gp.Acct1.Name == ACCT1 {
		keyRegex := `^[A-Za-z0-9+/]+={0,2}$`
		matched, err := regexp.MatchString(keyRegex, gp.Acct1.Details.Address)
		if err != nil {
			fmt.Println("Error with regex:", whereami.WhereAmI(), err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("Invalid PubKey match.")
			fmt.Println("Given PubKey:", gp.Acct1.Details.Address)
			os.Exit(1)
		}
	}

	if gp.Acct2.Name == ACCT2 {
		keyRegex := `^[A-Za-z0-9+/]+={0,2}$`
		matched, err := regexp.MatchString(keyRegex, gp.Acct2.Details.Address)
		if err != nil {
			fmt.Println("Error with regex:", whereami.WhereAmI(), err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("Invalid PubKey match.")
			fmt.Println("Given PubKey:", gp.Acct2.Details.Address)
			os.Exit(1)
		}
	}

	if gp.ChainID != settings.ChainID {
		fmt.Println("Error: chainID is not correct, required: ", settings.ChainID)
		os.Exit(1)
	}

	regex := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{6,9}Z$` // valid time format: 2025-01-17T02:42:28.062004646Z
	matched, err := regexp.MatchString(regex, gp.CreatedTime)
	if err != nil {
		fmt.Println("Error with regex:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Given createdTime:", gp.CreatedTime)
		fmt.Println("Error: createdTime is not in the correct format")
		os.Exit(1)
	}
	fmt.Println("Validation OK")
}
