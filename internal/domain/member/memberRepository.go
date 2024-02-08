package member

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/utils/jwt"
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

func FindByNameAndEmail(name string, email string) (models.Member, error) {
	var result models.Member

	query := "SELECT * FROM member WHERE name=? AND email=?"
	var address sql.NullString
	err := db.MyDb.QueryRow(query, name, email).Scan(&result.ID, &result.Name, &result.Email, &address, &result.RegDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, fmt.Errorf("member not found")
		}
		return result, err
	}
	if address.Valid {
		result.Address = address.String
	} else {
		result.Address = ""
	}

	log.Println(result.ID, result.Name, result.Email, result.Address, result.RegDate)
	return result, nil
}

func InsertMember(member *models.Member) error {
	tx, err := db.MyDb.Begin()
	queryMember := "INSERT INTO member (name, email) VALUES (?, ?)"
	result, err := tx.Exec(queryMember, member.Name, member.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	memberID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	queryAuthority := "INSERT INTO authorities VALUES (?,?)"
	result, err = tx.Exec(queryAuthority, memberID, "ROLE_GUEST")
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func SaveAddress(id int, addr string) (string, error) {
	tx, err := db.MyDb.Begin()
	UpdateAddress := "UPDATE member SET address = ? WHERE id = ?"
	_, err = tx.Exec(UpdateAddress, addr, id)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to update address: %v", err)
	}

	token, err := jwt.AddressTokenProvider(id, addr)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to make addrToken: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return token, err
	}
	return token, nil
}

func DeleteMember(id int) error {
	DeleteMemberQuery := "Delete from member where id = ?"
	_, err := db.MyDb.Exec(DeleteMemberQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func findAllMembers(page int) ([]models.Member, error) {
	pageSize := 20
	offset := (page - 1) * pageSize

	query := "SELECT * FROM member LIMIT ? OFFSET ?"
	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member

	for rows.Next() {
		member := models.Member{}
		s := reflect.ValueOf(&member).Elem()
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
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return members, nil
}

func findMemberById(id int) (models.Member, error) {
	member := models.Member{}
	s := reflect.ValueOf(&member).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}

	query := "SELECT * FROM member WHERE id=?"
	err := db.MyDb.QueryRow(query, id).Scan(columns...)
	if err != nil {
		if err == sql.ErrNoRows {
			return member, fmt.Errorf("member not found")
		}
		return member, err
	}
	return member, nil
}
