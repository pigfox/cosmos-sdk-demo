package main

type RegularAccount struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	PubKey  string `json:"pubkey"`
	Type    string `json:"type"`
}
