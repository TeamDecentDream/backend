package todo

import (
	"backend/internal/db"
	"backend/internal/models"
	"time"
)

func saveTodo(todo *models.Todo, id int) error {
	query := "insert into todo(author_id,contents) VALUES (?,?)"

	_, err := db.MyDb.Exec(query, id, todo.Contents)
	if err != nil {
		return err
	}
	return nil
}

func updateTodo(todo *models.Todo, id int) error {
	loc, err := time.LoadLocation("Asia/Seoul")
	currentTime := time.Now().In(loc)

	query := "update todo set contents=?, update_date=?, state=? where id =? and author_id=?"

	_, err = db.MyDb.Exec(query, todo.Contents, currentTime, todo.State, todo.Id, id)
	if err != nil {
		return err
	}
	return nil
}

func deleteTodo(id int, authorId int) error {
	query := "delete from todo where id=? and author_id = ?"
	_, err := db.MyDb.Exec(query, id, authorId)
	if err != nil {
		return err
	}
	return nil
}

func getTodos(dateRange *models.DateRange, id int) ([]models.Todo, error) {
	end := dateRange.End.AddDate(0, 0, 1)

	query := "SELECT * FROM todo WHERE reg_date >= ? AND reg_date < ? AND author_id=?"
	rows, err := db.MyDb.Query(query, dateRange.Start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.AuthorId, &todo.Contents, &todo.State, &todo.RegDate, &todo.UpdateDate)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
