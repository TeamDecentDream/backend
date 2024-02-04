package models

import "time"

type Notification struct {
	ID         int       `json:"id"`
	Contents   string    `json:"contents"`
	AuthorID   int       `json:"author_id"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}
