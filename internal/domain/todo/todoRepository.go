package todo

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"time"
)

func saveTodo(todo *models.Todo, id int) error {
	query := "insert into todo(author_id,contents,reg_date) VALUES (?,?,?)"

	_, err := db.MyDb.Exec(query, id, todo.Contents, todo.RegDate)
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

func getTodos(dateRange string, id int) ([]models.Todo, error) {
	layout := "2006-01"
	start, err := time.Parse(layout, dateRange)
	if err != nil {
		return nil, err
	}
	end := start.AddDate(0, 1, 0)

	query := "SELECT * FROM todo WHERE reg_date >= ? AND reg_date < ? AND author_id=?"
	rows, err := db.MyDb.Query(query, start, end, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var _update_date sql.NullTime
		err := rows.Scan(&todo.Id, &todo.AuthorId, &todo.Contents, &todo.State, &todo.RegDate, &_update_date)
		if err != nil {
			return nil, err
		}
		if _update_date.Valid {
			todo.UpdateDate = _update_date.Time
		}
		todos = append(todos, todo)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
