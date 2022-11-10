package domain

import "time"

type Transaction struct {
	Id         int64     `json:"id" db:"id"`
	SenderId   int64     `json:"senderId" db:"sender_id"`
	ReceiverId int64     `json:"receiverId" db:"receiver_id"`
	Amount     int64     `json:"amount" db:"amount"`
	Status     string    `json:"status" db:"status"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}
