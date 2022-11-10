package domain

type Client struct {
	Id      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int64  `json:"balance" db:"balance"`
}
