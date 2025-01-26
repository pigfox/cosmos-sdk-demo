package main

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Details struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Public  PubKey `json:"pub_key"`
}

type Account struct {
	Name    string  `json:"key_name"`
	Details Details `json:"account_key"`
}

func newAccount(keyName string) Account {
	return Account{
		Name: keyName,
	}
}

type GenesisParams struct {
	CreatedTime     string  `json:"created_time"`
	ChainID         string  `json:"chain_id"`
	Acct1           Account `json:"acct1"`
	Acct2           Account `json:"acct2"`
	Validator       Account `json:"validator"`
	Amount          string  `json:"amount"`
	ValidatorAmount string  `json:"validator_amount"`
	SupplyAmount    string  `json:"supply_amount"`
}

type ValidatorFile struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PubKey      string `json:"pubkey"`
	KeyringPath string `json:"keyring_path"`
	HomeDir     string `json:"home_dir"`
}

type Settings struct {
	KeyringBackend  string `json:"keyring_backend"`
	KeyName         string `json:"key_name"`
	ChainID         string `json:"chain_id"`
	Moniker         string `json:"moniker"`
	ValidatorPath   string `json:"validator_path"`
	Fees            string `json:"fees"`
	SupplyAmount    string `json:"supply_amount"`
	ValidatorAmount string `json:"validator_amount"`
	GenesisPath     string `json:"genesis_path"`
	AppHomeDir      string `json:"app_home_dir"`
	NodeURL         string `json:"node_url"`
}

func newSettings() Settings {
	appHomeDir := getHomeDir() + "/.simd"
	return Settings{
		KeyringBackend:  "test",
		KeyName:         "my-key",
		ChainID:         "my-chain",
		Moniker:         "pigfox",
		ValidatorPath:   appHomeDir + "/config/validator.json",
		Fees:            "10000",
		SupplyAmount:    "900000000",
		ValidatorAmount: "100000000",
		GenesisPath:     appHomeDir + "/config/genesis.json",
		AppHomeDir:      appHomeDir,
		NodeURL:         "http://localhost:26657",
	}
}
