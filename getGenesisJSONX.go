package main

/*
Docker generated file structure
*/
func getGenesisJSONX(gp GenesisParams) string {
	return `{
  "app_name": "simd",
  "app_version": "0.52.0-rc.1",
  "genesis_time": "` + gp.CreatedTime + `",
  "chain_id": "` + gp.ChainID + `",
  "initial_height": 1,
  "app_hash": null,
  "app_state": {
    "accounts": {
      "accounts": [],
      "init_account_msgs": []
    },
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": []
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [],
      "supply": [],
      "denom_metadata": [],
      "send_enabled": []
    },
    "circuit": {
      "account_permissions": [],
      "disabled_type_urls": []
    },
    "consensus": null,
    "counter": {},
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": [],
        "decimal_pool": []
      },
      "delegator_withdraw_infos": [],
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "epochs": {
      "epochs": [
        {
          "identifier": "day",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "86400s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        },
        {
          "identifier": "hour",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "3600s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        },
        {
          "identifier": "minute",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "60s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        },
        {
          "identifier": "week",
          "start_time": "0001-01-01T00:00:00Z",
          "duration": "604800s",
          "current_epoch": "0",
          "current_epoch_start_time": "0001-01-01T00:00:00Z",
          "epoch_counting_started": false,
          "current_epoch_start_height": "0"
        }
      ]
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
        "min_deposit_ratio": "0.010000000000000000",
        "proposal_cancel_max_period": "0.500000000000000000",
        "optimistic_authorized_addresses": [],
        "optimistic_rejected_threshold": "0.100000000000000000",
        "yes_quorum": "0.000000000000000000",
        "expedited_quorum": "0.500000000000000000",
        "proposal_execution_gas": "10000000"
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
        "annual_provisions": "0.000000000000000000",
        "data": null
      },
      "params": {
        "mint_denom": "stake",
        "inflation_rate_change": "0.130000000000000000",
        "inflation_max": "0.050000000000000000",
        "inflation_min": "0.000000000000000000",
        "goal_bonded": "0.670000000000000000",
        "blocks_per_year": "6311520",
        "max_supply": "0"
      }
    },
    "nft": {
      "classes": [],
      "entries": []
    },
    "protocolpool": {
      "continuous_fund": [],
      "budget": [],
      "last_balance": {
        "amount": []
      },
      "distributions": [],
      "params": {
        "enabled_distribution_denoms": [
          "stake"
        ]
      }
    },
    "runtime": {},
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
        "historical_entries": 0,
        "bond_denom": "stake",
        "min_commission_rate": "0.000000000000000000",
        "key_rotation_fee": {
          "denom": "stake",
          "amount": "1000000"
        }
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false,
      "rotation_index_records": [],
      "rotation_history": [],
      "rotation_queue": []
    },
    "upgrade": {},
    "validate": {},
    "vesting": {}
  },
  "consensus": {
    "params": {
      "block": {
        "max_bytes": "4194304",
        "max_gas": "10000000"
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
      "synchrony": {
        "precision": "505000000",
        "message_delay": "15000000000"
      },
      "feature": {
        "vote_extensions_enable_height": "0",
        "pbts_enable_height": "0"
      }
    }
  }
}`
}
