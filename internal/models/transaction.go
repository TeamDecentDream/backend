package models

import "time"

type Transaction struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Amount  int       `json:"amount"`
	Client  string    `json:"client"`
	SellBuy int       `json:"sell_buy"`
	RegDate time.Time `json:"reg_date"`
}
