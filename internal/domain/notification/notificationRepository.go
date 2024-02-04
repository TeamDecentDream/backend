package notification

import (
	"backend/internal/db"
	"backend/internal/models"
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
