package main

import (
	"errors"
)

type EventType string

const (
	Deposit  EventType = "deposit"
	Withdraw EventType = "withdraw"
	Transfer EventType = "transfer"
)

type Account struct {
	AccountID string  `json:"id"`
	Balance   float64 `json:"balance"`
}

type Event struct {
	Type        EventType `json:"type"`
	Destination string    `json:"destination"`
	Amount      float64   `json:"amount"`
}

var currentAccounts []Account = nil
var currentEvents []Event = nil

func ResetDB() {
	currentAccounts = []Account{}
	currentEvents = []Event{}
}

func CreateNewAccount(a Account) {
	currentAccounts = append(currentAccounts, a)
}

func CreateNewEvent(e Event) {
	currentEvents = append(currentEvents, e)
}

func GetAccount(id string) (Account, error) {
	found := linearSearchAccountsById(id)
	if found != -1 {
		return currentAccounts[found], nil
	}
	return Account{}, errors.New("cannot find account")
}

func linearSearchAccountsById(id string) int {
	for i := 0; i < len(currentAccounts); i++ {
		if currentAccounts[i].AccountID == id {
			return i
		}
	}
	return -1
}
