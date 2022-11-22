package models

type AccountSession struct {
	AccountID int64  `json:"account_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
