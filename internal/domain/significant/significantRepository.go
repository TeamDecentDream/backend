package significant

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"reflect"
)

func saveSignificant(input *models.Significant) error {
	query := "Insert into significant(contents, author_id) VALUE (?,?)"
	_, err := db.MyDb.Exec(query, input.Contents, input.AuthorID)
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

func getSignificants(page int) ([]models.SignificantOutput, error) {
	pageSize := 20
	offset := (page - 1) * pageSize

	query := "SELECT s.id, s.contents, m.name, s.warn ,s.reg_date,s.update_date FROM significant s join nextfarm.member m on m.id = s.author_id order by s.update_date desc LIMIT ? OFFSET ?"
	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var significants []models.SignificantOutput

	for rows.Next() {
		s := models.SignificantOutput{}
		err := rows.Scan(&s.ID, &s.Contents, &s.AuthorID, &s.Warn, &s.RegDate, &s.UpdateDate)
		if err != nil {
			return nil, err
		}
		significants = append(significants, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return significants, nil
}

func findSignificantById(id int) (models.Significant, error) {
	significant := models.Significant{}
	s := reflect.ValueOf(&significant).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	query := "SELECT * FROM significant WHERE id=?"
	err := db.MyDb.QueryRow(query, id).Scan(columns...)
	if err != nil {
		if err == sql.ErrNoRows {
			return significant, fmt.Errorf("member not found")
		}
		return significant, err
	}
	return significant, nil
}

func updateSignificant(updateInfo *models.Significant) error {
	query := "UPDATE significant SET contents=?, update_date=NOW(), author_id=?, warn=? WHERE id=?"

	_, err := db.MyDb.Exec(query, updateInfo.Contents, updateInfo.AuthorID, updateInfo.Warn, updateInfo.ID)
	if err != nil {
		return err
	}
	return err
}

func DeleteSignificant(sid int, id int) error {
	query := "Delete from significant where id=? AND author_id=?"
	_, err := db.MyDb.Exec(query, sid, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSignificantByAdmin(sid int) error {
	query := "Delete from significant where id=?"
	_, err := db.MyDb.Exec(query, sid)
	if err != nil {
		return err
	}
	return nil
}
