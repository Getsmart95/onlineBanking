package models

type Client struct {
	ID         int
	Name       string
	Surname    string
	Login      string
	Password   string
	Age        int
	Gender     string
	Phone      int
	Status 	   boolean
	VerifiedAt datetime
}

type (
	ClientList struct {
		Clients []Client
	}
)
