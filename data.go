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
	Origin      string    `json:"origin"`
}

type EventDepositResponse struct {
	Destination Account `json:"destination"`
}

type EventWithdrawResponse struct {
	Origin Account `json:"origin"`
}

type EventTransferResponse struct {
	Destination Account `json:"destination"`
	Origin      Account `json:"origin"`
}

var currentAccounts []Account = nil
var currentEvents []Event = nil

func ResetDB() {
	currentAccounts = []Account{}
	currentEvents = []Event{}
}

func CreateNewAccount(a Account) *Account {
	currentAccounts = append(currentAccounts, a)
	return &currentAccounts[len(currentAccounts)-1]
}

func CreateNewEvent(e Event) {
	currentEvents = append(currentEvents, e)
}

func CreateNewDepositEvent(e Event, acc *Account) {
	CreateNewEvent(e)
	DepositToAccount(acc, e.Amount)
}

func CreateNewWithdrawEvent(e Event, acc *Account) {
	CreateNewEvent(e)
	WithdrawFromAccount(acc, e.Amount)
}

func CreateNewTransferEvent(e Event, origin *Account, destination *Account) {
	CreateNewEvent(e)
	WithdrawFromAccount(origin, e.Amount)
	DepositToAccount(destination, e.Amount)
}

func WithdrawFromAccount(acc *Account, value float64) {
	acc.Balance -= value
}

func DepositToAccount(acc *Account, value float64) {
	acc.Balance += value
}

func GetAccount(id string) (*Account, error) {
	found := linearSearchAccountsById(id)
	if found != nil {
		return found, nil
	}
	return nil, errors.New("cannot find account")
}

func linearSearchAccountsById(id string) *Account {
	for i := 0; i < len(currentAccounts); i++ {
		if currentAccounts[i].AccountID == id {
			return &currentAccounts[i]
		}
	}
	return nil
}
