package significant

import (
	"backend/internal/db"
	"backend/internal/models"
)

func saveSignificant(input *models.Significant) error {
	query := "Insert into significant(title, contents, author_id) VALUE (?,?,?)"
	_, err := db.MyDb.Exec(query, input.Title, input.Contents, input.AuthorID)
	if err != nil {
		return err
	}
	return err
}
