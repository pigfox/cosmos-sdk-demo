package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/jimlawless/whereami"
)

/*
// Struct to represent the expected PubKeyData JSON structure

	type PubKeyData struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	}
*/
func addGenesis(accountAddress, validatorAddress string, validatorPubkey PubKey) {
	fmt.Println("addGenesis: Create the genesis file")

	// Define the target path for the genesis file
	genesisFile := settings.GenesisPath

	created := time.Now().UTC().Format(time.RFC3339Nano)
	gp := GenesisParams{
		CreatedTime:      created,
		ChainID:          settings.ChainID,
		Address:          accountAddress,
		PubKEY:           validatorPubkey,
		ValidatorAddress: validatorAddress,
	}
	fmt.Println("accountAddress", accountAddress)
	fmt.Printf("+%v\n", gp)

	genesisJson := getGenesisJSON(gp)

	// Write the updated data to the target genesis file location
	err := os.WriteFile(genesisFile, []byte(genesisJson), 0644)
	if err != nil {
		fmt.Println("Error: Failed to write updated genesis file:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	// JSON is valid
	fmt.Println("Genesis file has been created and is valid. Written to:", genesisFile)
}

func getGenesisJSON(gp GenesisParams) string {
	fmt.Println("address:", gp.Address)
	regex := `^cosmos1[a-z0-9]{38}$`
	// Validate the address
	matched, err := regexp.MatchString(regex, gp.Address)
	if err != nil {
		fmt.Println("Error with regex:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Error: address is not in the correct format", whereami.WhereAmI())
		fmt.Println("address:", gp.Address)
		os.Exit(1)
	}

	regex = `^cosmosvaloper1[a-z0-9]{38}$`
	matched, err = regexp.MatchString(regex, gp.ValidatorAddress)
	if err != nil {
		fmt.Println("Error with regex:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Error: validator address is not in the correct format", whereami.WhereAmI())
		fmt.Println("validatorAddress:", gp.ValidatorAddress, "length", len(gp.ValidatorAddress))
		os.Exit(1)
	}

	fmt.Printf("+%v", gp)
	fmt.Println("gp.PubKEY.Key", gp.PubKEY.Key)

	// Regex to validate the "key" field as base64 encoded string with padding
	keyRegex := `^[a-zA-Z0-9+/]{43}=$`
	matched, err = regexp.MatchString(keyRegex, gp.PubKEY.Key)
	if err != nil {
		fmt.Println("Error with regex:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Valid PubKey format.")
		os.Exit(1)
	}

	if gp.ChainID != settings.ChainID {
		fmt.Println("Error: chainID is not correct, required: ", settings.ChainID)
		os.Exit(1)
	}

	regex = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{8,9}Z$` // vaild time format: 2025-01-17T02:42:28.062004646Z
	matched, err = regexp.MatchString(regex, gp.CreatedTime)
	if err != nil {
		fmt.Println("Error with regex:", whereami.WhereAmI(), err)
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Given createdTime:", gp.CreatedTime)
		fmt.Println("Error: createdTime is not in the correct format")
		os.Exit(1)
	}

	return `{
  "app_name": "simd",
  "app_version": "0.50.6",
  "genesis_time": "` + gp.CreatedTime + `",
  "chain_id": "` + gp.ChainID + `",
  "initial_height": 1,
  "app_hash": null,
  "app_state": {
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "` + gp.Address + `",
          "pub_key": {
            "@type": "` + gp.PubKEY.Type + `",
            "key": "` + gp.PubKEY.Key + `"
          },
          "account_number": "0",
          "sequence": "0"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "` + gp.Address + `",
          "coins": [
            {
              "denom": "stake",
              "amount": "100000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "stake",
          "amount": "100000000"
        }
      ],
      "denom_metadata": [],
      "send_enabled": []
    },
    "circuit": {
      "account_permissions": [],
      "disabled_type_urls": []
    },
    "consensus": null,
    "crisis": {
      "constant_fee": {
        "denom": "stake",
        "amount": "1000"
      }
    },
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "evidence": {
      "evidence": []
    },
    "feegrant": {
      "allowances": []
    },
    "genutil": {
      "gen_txs": []
    },
    "gov": {
      "starting_proposal_id": "1",
      "deposits": [],
      "votes": [],
      "proposals": [],
      "deposit_params": null,
      "voting_params": null,
      "tally_params": null,
      "params": {
        "min_deposit": [
          {
            "denom": "stake",
            "amount": "10000000"
          }
        ],
        "max_deposit_period": "172800s",
        "voting_period": "172800s",
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000",
        "min_initial_deposit_ratio": "0.000000000000000000",
        "proposal_cancel_ratio": "0.500000000000000000",
        "proposal_cancel_dest": "",
        "expedited_voting_period": "86400s",
        "expedited_threshold": "0.667000000000000000",
        "expedited_min_deposit": [
          {
            "denom": "stake",
            "amount": "50000000"
          }
        ],
        "burn_vote_quorum": false,
        "burn_proposal_deposit_prevote": false,
        "burn_vote_veto": true,
        "min_deposit_ratio": "0.010000000000000000"
      },
      "constitution": ""
    },
    "group": {
      "group_seq": "0",
      "groups": [],
      "group_members": [],
      "group_policy_seq": "0",
      "group_policies": [],
      "proposal_seq": "0",
      "proposals": [],
      "votes": []
    },
    "mint": {
      "minter": {
        "inflation": "0.130000000000000000",
        "annual_provisions": "0.000000000000000000"
      },
      "params": {
        "mint_denom": "stake",
        "inflation_rate_change": "0.130000000000000000",
        "inflation_max": "0.200000000000000000",
        "inflation_min": "0.070000000000000000",
        "goal_bonded": "0.670000000000000000",
        "blocks_per_year": "6311520"
      }
    },
    "nft": {
      "classes": [],
      "entries": []
    },
    "params": null,
    "runtime": null,
    "slashing": {
      "params": {
        "signed_blocks_window": "100",
        "min_signed_per_window": "0.500000000000000000",
        "downtime_jail_duration": "600s",
        "slash_fraction_double_sign": "0.050000000000000000",
        "slash_fraction_downtime": "0.010000000000000000"
      },
      "signing_infos": [],
      "missed_blocks": []
    },
    "staking": {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "stake",
        "min_commission_rate": "0.000000000000000000"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [
  {
    "operator_address": "cosmosvaloper1abcd...",
    "consensus_pubkey": {
      "@type": "/cosmos.crypto.ed25519.PubKey",
      "key": "4ZG5UqZ..."
    },
    "status": "BOND_STATUS_BONDED",
    "tokens": "100000000",
    "delegator_shares": "100000000.000000000000000000",
    "description": {
      "moniker": "MyValidator",
      "identity": "",
      "website": "",
      "details": ""
    },
    "unbonding_height": "0",
    "unbonding_time": "1970-01-01T00:00:00Z",
    "commission": {
      "commission_rates": {
        "rate": "0.100000000000000000",
        "max_rate": "0.200000000000000000",
        "max_change_rate": "0.010000000000000000"
      },
      "update_time": "1970-01-01T00:00:00Z"
    },
    "min_self_delegation": "1"
  }
],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false
    },
    "upgrade": {},
    "vesting": {}
  },
  "consensus": {
    "params": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      },
      "version": {
        "app": "0"
      },
      "abci": {
        "vote_extensions_enable_height": "0"
      }
    }
  }
}`
}
