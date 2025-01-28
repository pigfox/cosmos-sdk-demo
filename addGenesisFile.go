package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jimlawless/whereami"
)

func addGenesisFile(acct1, acct2, validator Account) {
	fmt.Println("addGenesisFile()")

	// Define the target path for the genesis file
	genesisFile := settings.GenesisPath
	validator.Details.Address = strings.TrimSpace(validator.Details.Address)

	created := time.Now().UTC().Format(time.RFC3339Nano)
	gp := GenesisParams{
		CreatedTime:     created,
		ChainID:         settings.ChainID,
		Acct1:           acct1,
		Acct2:           acct2,
		Validator:       validator,
		ValidatorAmount: settings.ValidatorAmount,
		SupplyAmount:    settings.SupplyAmount,
	}
	accountValidation(gp)
	genesisJson := getGenesisJSONX(gp)
	//genesisJson := getGenesisJSON0(gp)

	// Write the updated data to the target genesis file location
	err := os.WriteFile(genesisFile, []byte(genesisJson), 0644)
	if err != nil {
		fmt.Println("Error: Failed to write updated genesis file:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	validateCmd := []string{"genesis", "validate", settings.GenesisPath}
	output, err := simdCmd(validateCmd)

	if err != nil {
		fmt.Println("Error: Failed to fetch regular account address:", err)
		fmt.Println("Command Output:", string(output))
		os.Exit(1)
	} else {
		fmt.Println(string(output))
	}

	cmd := exec.Command("ls", "-l", settings.GenesisPath)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(out))

	cmd = exec.Command("cat", settings.GenesisPath)
	out, err = cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(out))

}
