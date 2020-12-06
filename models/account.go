package models

import "time"

type Account struct {
	ID            int
	ClientId      int
	AccountNumber int
	Balance       int
	Status        bool
	CardNumber    string
	LimitTransfer int
	LimitPayment  int
	CreatedAt     time.Time
	UntilAt       time.Time
}

type AccountList struct {
	Accounts []Account
}

type AccountWithUserName struct {
	Account Account
	Client  Client
}

type AccountForUser struct {
	ID            int64
	Name          string
	AccountNumber int64
	Balance       int64
	Locked        bool
}
