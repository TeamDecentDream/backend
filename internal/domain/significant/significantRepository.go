package significant

import (
	"backend/internal/db"
	"backend/internal/models"
	"reflect"
)

func saveSignificant(input *models.Significant) error {
	query := "Insert into significant(title, contents, author_id) VALUE (?,?,?)"
	_, err := db.MyDb.Exec(query, input.Title, input.Contents, input.AuthorID)
	if err != nil {
		return err
	}
	return err
}

func GetSignificantCount() (int, error) {
	query := "select count(*) from significant"
	var count int
	err := db.MyDb.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, err
}

func getSignificants(page int) ([]models.Significant, error) {
	pageSize := 20
	offset := (page - 1) * pageSize

	query := "SELECT * FROM significant LIMIT ? OFFSET ?"
	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var significants []models.Significant

	for rows.Next() {
		significant := models.Significant{}
		s := reflect.ValueOf(&significant).Elem()
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
		significants = append(significants, significant)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return significants, nil
}
