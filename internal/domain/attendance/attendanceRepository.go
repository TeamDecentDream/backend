package attendance

import (
	"backend/internal/db"
	"backend/internal/models"
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

	var enterLog models.Attendance
	var leavTime sql.NullTime
	err = db.MyDb.QueryRow(query, memberId, today).Scan(&enterLog.Id, &enterLog.MemberId, &enterLog.EnterTime, &leavTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return &enterLog.EnterTime, sql.ErrNoRows
		}
		return &enterLog.EnterTime, err
	}

	return &enterLog.EnterTime, nil
}

func enter(memberId int) (*time.Time, error) {
	var timelog time.Time
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO attendance (member_id) VALUES (?)"
	_, err = db.MyDb.Exec(query, memberId)
	if err != nil {
		return nil, err
	}
	timelog = time.Now().In(loc)
	return &timelog, nil
}

func leave(memberId int) error {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return err
	}

	now := time.Now().In(loc)
	today := now.Format("2006-01-02")

	query := "UPDATE attendance SET leave_time = ? WHERE member_id = ? AND DATE(enter_time) = ? AND leave_time IS NULL"

	_, err = db.MyDb.Exec(query, now, memberId, today)
	if err != nil {
		return err
	}
	return nil
}

func GetWorkTimeLog(memberId int) ([]models.Attendance, error) {
	query := "select * from attendance where member_id=? order by enter_time desc limit 5"

	rows, err := db.MyDb.Query(query, memberId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var attendance models.Attendance
		var leaveTime sql.NullTime
		err := rows.Scan(&attendance.Id, &attendance.MemberId, &attendance.EnterTime, &leaveTime)
		if err != nil {
			return nil, err
		}
		if leaveTime.Valid {
			attendance.LeaveTime = leaveTime.Time
		}
		attendances = append(attendances, attendance)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}
