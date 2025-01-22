package main

type RegularAccount struct {
	Address              string               `json:"address"`
	Name                 string               `json:"name"`
	RegularAccountPubKey RegularAccountPubKey `json:"pubkey"`
	Type                 string               `json:"type"`
}

type RegularAccountPubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type PubKey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type ValidatorKeyData struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
}

type GenesisParams struct {
	CreatedTime      string `json:"created_time"`
	ChainID          string `json:"chain_id"`
	Address          string `json:"address"`
	PubKEY           string `json:"pub_key"`
	ValidatorAddress string `json:"validator_address"`
}

type Settings struct {
	KeyringBackend string `json:"keyring_backend"`
	KeyName        string `json:"key_name"`
	ChainID        string `json:"chain_id"`
	ValidatorName  string `json:"validator_name"`
	ValidatorPath  string `json:"validator_path"`
	Fees           string `json:"fees"`
	Amount         string `json:"amount"`
	GenesisPath    string `json:"genesis_path"`
	AppHomeDir     string `json:"app_home_dir"`
}

func newSettings() Settings {
	appHomeDir := getHomeDir() + "/.simapp"
	return Settings{
		KeyringBackend: "test",
		KeyName:        "my-key",
		ChainID:        "my-chain",
		ValidatorName:  "pigfox",
		ValidatorPath:  appHomeDir + "/validator.json",
		Fees:           "10000stake",
		Amount:         "500000stake",
		GenesisPath:    appHomeDir + "/config/genesis.json",
		AppHomeDir:     appHomeDir,
	}
}
