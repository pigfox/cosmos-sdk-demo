package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting setup...")
	validatorName := "my-validator"
	chainID := "123456789"
	source := "./genesis.json"
	destination := getHomeDir() + "/.simapp/config/genesis.json"
	step0()
	step1()
	step2()
	step3(validatorName)
	validatorAddress := step4(validatorName)
	validatorPubkey := step5()
	//validatorOperatorAddress := step6(validatorAddress)
	step7(validatorAddress, validatorPubkey, chainID)
	step8(source, destination)
	step9(destination)
	step10()
}
