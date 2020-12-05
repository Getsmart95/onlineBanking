package models

type Service struct{
	ID int
	Name string
}

type ServiceList struct{
	Services []Service
}
