package models

type EvaluationInput struct {
	TextValue    string `json:"textValue"`
	WorkJournals []int  `json:"workJournals"`
	ID           int    `json:"id"`
}
