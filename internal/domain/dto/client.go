package dto

type CreateClientDto struct {
	Name    string `json:"name" required:"true"`
	Balance int64  `json:"balance" required:"true"`
}
