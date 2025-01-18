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
