package models

import "time"

type Todo struct {
	Id         int       `json:"id"`
	AuthorId   int       `json:"authorId"`
	Contents   string    `json:"contents"`
	State      int       `json:"state"`
	UpdateDate time.Time `json:"updateDate"`
	RegDate    time.Time `json:"regDate"`
}

type DateRange struct {
	Start time.Time
	End   time.Time
}
