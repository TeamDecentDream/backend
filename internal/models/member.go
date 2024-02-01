package models

import "time"

type Member struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	RegDate time.Time `json:"reg_date"`
}
