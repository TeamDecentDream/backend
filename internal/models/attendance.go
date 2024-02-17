package models

import "time"

type Attendance struct {
	Id        int       `json:"id"`
	MemberId  int       `json:"member_id"`
	EnterTime time.Time `json:"enter_time"`
	LeaveTime time.Time `json:"leave_time"`
}
