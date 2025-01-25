package main

import (
	"fmt"
	"github.com/jimlawless/whereami"
	"os"
	"regexp"
)

func accountValidation(gp GenesisParams) {
	if gp.Validator.KeyName == VALIDATOR {
		regex := `^cosmosvaloper1[a-z0-9]{38}$`
		matched, err := regexp.MatchString(regex, gp.Validator.AccountKey.Address)
		if err != nil {
			fmt.Println("Error with regex:", whereami.WhereAmI(), err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("Error: validator address is not in the correct format", whereami.WhereAmI())
			fmt.Println("validatorAddress:", gp.Validator.AccountKey.Address, "length", len(gp.Validator.AccountKey.Address))
			os.Exit(1)
		}
	}
	/*else {
		keyRegex := `^[A-Za-z0-9+/]+={0,2}$`
		matched, err := regexp.MatchString(keyRegex, account.AccountKey.Address)
		if err != nil {
			fmt.Println("Error with regex:", whereami.WhereAmI(), err)
			os.Exit(1)
		}

		if !matched {
			fmt.Println("Invalid PubKey match.")
			fmt.Println("Given PubKey:", account.AccountKey.Address)
			os.Exit(1)
		}
	}
	*/
	if gp.ChainID != settings.ChainID {
		fmt.Println("Error: chainID is not correct, required: ", settings.ChainID)
		os.Exit(1)
	}
}
