package models

import "time"

type Notification struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Contents   string    `json:"contents"`
	AuthorID   int       `json:"author_id"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}

type NotificationOutput struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Contents   string    `json:"contents"`
	Author     string    `json:"author_id"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}
