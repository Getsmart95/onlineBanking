package models

type ATM struct {
	ID        int
	Name      string
	Status    boolean
	CreatedAt datetime
}

type ATMList struct{
	ATMs []ATM
}