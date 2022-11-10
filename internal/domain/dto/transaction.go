package dto

type CreateTransactionDto struct {
	SenderId   int64 `json:"senderId"`
	ReceiverId int64 `json:"receiverId"`
	Amount     int64 `json:"amount"`
}
