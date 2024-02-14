package notification

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"errors"
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
func getNotification(page int) ([]models.NotificationOutput, error) {
	pageSize := 5
	offset := (page - 1) * pageSize

	query := "SELECT n.id, n.title, n.contents, m.name, n.reg_date, n.update_date  FROM notification n join nextfarm.member m on m.id = n.author_id ORDER BY n.update_date DESC LIMIT ? OFFSET ? "
	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.NotificationOutput

	for rows.Next() {
		notification := models.NotificationOutput{}
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
	query := "Insert into notification(title,contents, author_id) VALUE (?,?,?)"
	_, err := db.MyDb.Exec(query, input.Title, input.Contents, input.AuthorID)
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
		if errors.Is(err, sql.ErrNoRows) {
			return notification, fmt.Errorf("Notification not found")
		}
		return notification, err
	}
	return notification, nil
}

func findNotificationOutputById(id int) (models.NotificationOutput, error) {
	notification := models.NotificationOutput{}

	query := "SELECT notification.id, title, contents, m.name, notification.reg_date, update_date FROM notification left join nextfarm.member m on notification.author_id = m.id WHERE id=?"
	err := db.MyDb.QueryRow(query, id).Scan(&notification.ID, &notification.Title, &notification.Contents, &notification.Author, &notification.RegDate, &notification.UpdateDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return notification, fmt.Errorf("Notification not found")
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
