package models

import "time"

type Client struct {
	ID         int
	Name       string
	Surname    string
	Login      string
	Password   string
	Age        int
	Gender     string
	Phone      string
	Status 	   bool
	VerifiedAt time.Time
}

type (
	ClientList struct {
		Clients []Client
	}
)
