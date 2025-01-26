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
	validator.AccountKey.Address = strings.TrimSpace(validator.AccountKey.Address)

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
	genesisJson := getGenesisJSON(gp)

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

func getGenesisJSON(gp GenesisParams) string {
	return ""
	/*
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
			          "address": "` + gp.Acct1.AccountKey.Address + `",
			          "pub_key": {
			            "@type": "` + gp.Acct1.AccountKey.PubKey.Type + `",
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
			          "address": "` + gp.Acct1.AccountKey.Address + `",
			          "coins": [
			            {
			              "denom": "stake",
			              "amount": "` + gp.SupplyAmount + `"
			            }
			          ]
			        }
			      ],
			      "supply": [
			        {
			          "denom": "stake",
			          "amount": "` + gp.SupplyAmount + `"
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
			    "operator_address": "` + gp.ValidatorAddress + `",
			    "consensus_pubkey": {
			      "@type": "` + gp.PubKEY.Type + `",
			      "key": "` + gp.PubKEY.Key + `"
			    },
			    "status": "BOND_STATUS_BONDED",
			    "tokens": "` + gp.ValidatorAmount + `",
			    "delegator_shares": "100000000.000000000000000000",
			    "description": {
			      "moniker": "` + settings.Moniker + `",
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
	*/
}
