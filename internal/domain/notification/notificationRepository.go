package notification

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"reflect"
)

func GetNotificationCount() (int, error) {
	query := "select count(*) from notification"
	var count int
	err := db.MyDb.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, err
}
func getNotification(page int) ([]models.Notification, error) {
	pageSize := 20
	offset := (page - 1) * pageSize

	query := "SELECT * FROM notification LIMIT ? OFFSET ?"
	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		notification := models.Notification{}
		s := reflect.ValueOf(&notification).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}

func saveNotification(input *models.Notification) error {
	query := "Insert into notification(contents, author_id) VALUE (?,?)"
	_, err := db.MyDb.Exec(query, input.Contents, input.AuthorID)
	if err != nil {
		return err
	}
	return err
}

func findNotificationById(id int) (models.Notification, error) {
	notification := models.Notification{}
	s := reflect.ValueOf(&notification).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	query := "SELECT * FROM notification WHERE id=?"
	err := db.MyDb.QueryRow(query, id).Scan(columns...)
	if err != nil {
		if err == sql.ErrNoRows {
			return notification, fmt.Errorf("member not found")
		}
		return notification, err
	}
	return notification, nil
}

func updateNotification(updateInfo *models.Notification, id int) error {
	query := "UPDATE notification SET contents=?, update_date=NOW(), author_id=? WHERE id=?"

	_, err := db.MyDb.Exec(query, updateInfo.Contents, id)
	if err != nil {
		return err
	}
	return err
}

func DeleteNotification(sid int, id int) error {
	query := "Delete from notification where id=? AND author_id=?"
	_, err := db.MyDb.Exec(query, sid, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteNotificationByAdmin(sid int) error {
	query := "Delete from notification where id=?"
	_, err := db.MyDb.Exec(query, sid)
	if err != nil {
		return err
	}
	return nil
}
