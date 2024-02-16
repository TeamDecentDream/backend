package models

import "time"

type Significant struct {
	ID         int       `json:"id"`
	Contents   string    `json:"contents"`
	AuthorID   int       `json:"author_id"`
	Warn       int       `json:"warn"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}

type SignificantOutput struct {
	ID         int       `json:"id"`
	Contents   string    `json:"contents"`
	Warn       int       `json:"warn"`
	AuthorID   string    `json:"author_id"`
	RegDate    time.Time `json:"reg_date"`
	UpdateDate time.Time `json:"update_date"`
}
