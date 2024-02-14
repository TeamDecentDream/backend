package models

import "time"

type Significant struct {
	ID         int       `json:"id"`
	Contents   string    `json:"contents"`
	Grade      int       `json:"grade"`
	AuthorID   int       `json:"author_id"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}
