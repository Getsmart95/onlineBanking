package models

type History struct {
	ID          int
	SenderId    int
	RecipientId int
	Money       int
	Message     string
	ServiceId   int
	CreatedAt   datetime
}

type HistoryList struct{
	Histories []History
}

