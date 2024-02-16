package attendance

import (
	"backend/internal/db"
	"database/sql"
	"time"
)

func GetWorkState(memberId int) (*time.Time, error) {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, err
	}
	
	today := time.Now().In(loc).Format("2006-01-02")
	query := "SELECT * FROM attendance WHERE member_id = ? AND DATE(enter_time) = ? AND leave_time is null order by enter_time desc"

	var enterTime time.Time
	err = db.MyDb.QueryRow(query, memberId, today).Scan(&enterTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return &enterTime, sql.ErrNoRows
		}
		return &enterTime, err
	}

	return &enterTime, nil
}
