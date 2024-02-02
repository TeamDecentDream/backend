package member

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/utils/jwt"
	"database/sql"
	"fmt"
	"log"
)

func FindByNameAndEmail(name string, email string) (models.Member, error) {
	var result models.Member
	query := "SELECT * FROM member WHERE name=? AND email=?"
	err := db.MyDb.QueryRow(query, name, email).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, fmt.Errorf("member not found")
		}
		return result, err
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
