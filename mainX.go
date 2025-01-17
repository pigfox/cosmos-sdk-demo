package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jimlawless/whereami"
)

const (
	keyringBackend = "test"
	chainID        = "1234567890"
)

func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running %s: %s", name, out.String())
	}
	return out.String(), nil
}

func mainX() {
	fmt.Println("---------Current script name: blockchain_setup.go---------")

	// Step 0: Clean up existing configuration and keys
	fmt.Println("Removing existing configuration and keys...")
	_ = os.RemoveAll("./config")

	keyListOutput, err := runCommand("simd", "keys", "list", "--home", "./", "--keyring-backend", keyringBackend, "--output", "json")
	if err != nil {
		fmt.Println("No keys to delete.")
	} else {
		var keys []map[string]interface{}
		if err := json.Unmarshal([]byte(keyListOutput), &keys); err == nil {
			for _, key := range keys {
				if name, ok := key["name"].(string); ok {
					fmt.Printf("Deleting key: %s\n", name)
					_, _ = runCommand("simd", "keys", "delete", name, "--home", "./", "--keyring-backend", keyringBackend, "-y")
				}
			}
		}
	}

	// Step 1: Display the current simd version
	fmt.Println("Displaying simd version...")
	version, err := runCommand("simd", "version")
	if err != nil {
		fmt.Println("Error fetching simd version:", err, whereami.WhereAmI())
	} else {
		fmt.Println("simd version:", version)
	}

	// Step 2: Initialize the blockchain
	fmt.Println("Initializing the blockchain...")
	_, err = runCommand("simd", "init", "pigfox", "--chain-id", chainID, "--home", "./")
	if err != nil {
		fmt.Println("Error initializing blockchain:", err, whereami.WhereAmI())
		return
	}

	// Step 3: Create a new key and genesis account
	keyName := fmt.Sprintf("pigfox_%d", time.Now().Unix())
	fmt.Printf("Creating new key: %s\n", keyName)
	_, err = runCommand("simd", "keys", "add", keyName, "--home", "./", "--keyring-backend", keyringBackend)
	if err != nil {
		fmt.Println("Error creating new key:", err, whereami.WhereAmI())
		return
	}

	address, err := runCommand("simd", "keys", "show", keyName, "-a", "--home", "./", "--keyring-backend", keyringBackend)
	if err != nil {
		fmt.Println("Error fetching key address:", err, whereami.WhereAmI())
		return
	}
	address = strings.TrimSpace(address)
	_, err = runCommand("simd", "genesis", "add-genesis-account", address, "100000000stake", "--home", "./")
	if err != nil {
		fmt.Println("Error adding genesis account:", err, whereami.WhereAmI())
		return
	}

	// Step 4: Add validator to genesis
	// Fetching validator public key
	fmt.Println("Fetching validator public key...")
	validatorPubKey, err := runCommand("simd", "tendermint", "show-validator", "--home", "./")
	if err != nil {
		fmt.Println("Error fetching validator public key:", err, whereami.WhereAmI())
		return
	}
	validatorPubKey = strings.TrimSpace(validatorPubKey)

	// Converting public key to validator address
	fmt.Printf("Validator Public Key: %s\n", validatorPubKey)
	validatorAddress, err := runCommand("simd", "debug", "pubkey", validatorPubKey)
	if err != nil {
		fmt.Println("Error extracting validator address:", err, whereami.WhereAmI())
		return
	}
	validatorAddress = strings.TrimSpace(validatorAddress)
	fmt.Printf("Validator Address: %s\n", validatorAddress)

	// Step 5: Delegate tokens to validator
	fmt.Println("Delegating tokens to validator...")
	_, err = runCommand("simd", "tx", "staking", "delegate", validatorAddress, "500000stake", "--from", keyName, "--chain-id", chainID, "--home", "./", "--keyring-backend", keyringBackend, "--broadcast-mode", "sync", "--yes", "--fees", "10000stake")
	if err != nil {
		fmt.Println("Error delegating tokens:", err, whereami.WhereAmI())
		return
	}

	// Step 6: Create a second key and transfer tokens
	secondKeyName := fmt.Sprintf("receiver_%d", time.Now().Unix())
	fmt.Printf("Creating second key: %s\n", secondKeyName)
	_, err = runCommand("simd", "keys", "add", secondKeyName, "--home", "./", "--keyring-backend", keyringBackend)
	if err != nil {
		fmt.Println("Error creating second key:", err, whereami.WhereAmI())
		return
	}
	secondAddress, err := runCommand("simd", "keys", "show", secondKeyName, "-a", "--home", "./", "--keyring-backend", keyringBackend)
	if err != nil {
		fmt.Println("Error fetching second key address:", err, whereami.WhereAmI())
		return
	}
	secondAddress = strings.TrimSpace(secondAddress)
	_, err = runCommand("simd", "tx", "bank", "send", address, secondAddress, "100000stake", "--chain-id", chainID, "--home", "./", "--keyring-backend", keyringBackend, "--yes")
	if err != nil {
		fmt.Println("Error transferring tokens:", err, whereami.WhereAmI())
		return
	}

	// Step 7: Submit and vote on a governance proposal
	fmt.Println("Submitting governance proposal...")
	_, err = runCommand("simd", "tx", "gov", "submit-proposal", "text", "Increase block size limit", "--description", "A proposal to increase block size", "--deposit", "1000000stake", "--from", keyName, "--chain-id", chainID, "--home", "./", "--keyring-backend", keyringBackend, "--yes")
	if err != nil {
		fmt.Println("Error submitting governance proposal:", err, whereami.WhereAmI())
		return
	}

	fmt.Println("Fetching latest proposal ID...")
	proposalOutput, err := runCommand("simd", "query", "gov", "proposals", "--output", "json")
	if err != nil {
		fmt.Println("Error fetching proposals:", err, whereami.WhereAmI())
		return
	}
	var proposals map[string]interface{}
	if err := json.Unmarshal([]byte(proposalOutput), &proposals); err == nil {
		if proposalList, ok := proposals["proposals"].([]interface{}); ok && len(proposalList) > 0 {
			lastProposal := proposalList[len(proposalList)-1].(map[string]interface{})
			proposalID := lastProposal["proposal_id"].(string)
			fmt.Printf("Voting on proposal ID: %s\n", proposalID)
			_, err = runCommand("simd", "tx", "gov", "vote", proposalID, "yes", "--from", keyName, "--chain-id", chainID, "--home", "./", "--keyring-backend", keyringBackend, "--yes")
			if err != nil {
				fmt.Println("Error voting on proposal:", err, whereami.WhereAmI())
			}
		}
	}

	fmt.Println("Blockchain setup and transactions completed successfully!")
}
