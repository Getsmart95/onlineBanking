package models

type Account struct {
	ID            int
	ClientId      int
	AccountNumber int
	Balance       int
	Status        boolean
	CardNumber    int
	LimitTransfer int
	LimitPayment  int
	CreatedAt     datetime
	UntilAt       datetime
}

type AccountList struct{
	Accounts []Account
}

type AccountWithUserName struct {
	Account Account
	Client 	Client
}

type AccountForUser struct {
	ID int64
	Name string
	AccountNumber int64
	Balance int64
	Locked bool
}