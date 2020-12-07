package models

import "time"

type History struct {
	ID          int
	SenderId    int
	RecipientId int
	Money       int
	Message     string
	ServiceId   int
	CreatedAt   time.Time
}

type HistoryList struct {
	Histories []History
}
