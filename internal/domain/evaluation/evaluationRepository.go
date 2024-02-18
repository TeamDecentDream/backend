package evaluation

import (
	"backend/internal/db"
	"backend/internal/models"
)

func saveEvaluation(input *models.EvaluationInput) error {
	query := "insert into evaluate(member_id, q1, q2, q3, q4, q5, q6, note) values (?,?,?,?,?,?,?,?)"
	_, err := db.MyDb.Exec(query, input.ID, input.WorkJournals[0], input.WorkJournals[1], input.WorkJournals[2], input.WorkJournals[3], input.WorkJournals[4], input.WorkJournals[5], input.TextValue)
	if err != nil {
		return err
	}
	return nil
}
