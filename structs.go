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

type Validator struct {
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
}

type Settings struct {
	KeyringBackend string `json:"keyring_backend"`
	KeyName        string `json:"key_name"`
	ChainID        string `json:"chain_id"`
	ValidatorName  string `json:"validator_name"`
	Fees           string `json:"fees"`
	Amount         string `json:"amount"`
	GenesisPath    string `json:"genesis_path"`
	AppHomeDir     string `json:"app_home_dir"`
}
