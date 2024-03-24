package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"miguel": {
		AuthToken: "123ABC",
		Username: "miguel",
	},
	"Jhon": {
		AuthToken: "456DEF",
		Username: "jhon",

	},
	"alysa": {
		AuthToken: "789GHI",
		Username: "alysa",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"miguel": {
		Coins: 200,
		Username: "miguel",
	},
	"Jhon": {
		Coins: 400,
		Username: "jhon",

	},
	"alysa": {
		Coins: 150,
		Username: "alysa",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB Call
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil

	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}