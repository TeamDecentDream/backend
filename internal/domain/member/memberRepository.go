package member

import (
	"backend/internal/db"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func FindByNameAndEmail(name string, email string) (models.Member, error) {
	var result models.Member

	query := "SELECT * FROM member left outer join nextfarm.authorities a on member.id = a.member_id WHERE name=? AND email=?"

	rows, err := db.MyDb.Query(query, name, email)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	authorities := make([]models.Authority, 0)
	found := false

	for rows.Next() {
		found = true
		var memberID int
		var role sql.NullString
		var address sql.NullString
		err := rows.Scan(&result.ID, &result.Name, &result.Email, &address, &result.RegDate, &memberID, &role)
		if err != nil {
			return result, err
		}

		if address.Valid {
			result.Address = address.String
		}

		if role.Valid {
			log.Println(role.String)
			authority := models.Authority{MemberId: memberID, Role: role.String}
			authorities = append(authorities, authority)
		}
	}
	if !found {
		return result, fmt.Errorf("member not found")
	}

	result.Authorities = authorities

	log.Println(result.ID, result.Name, result.Email, result.Address, result.RegDate, result.Authorities)
	return result, nil
}

func InsertMember(member *models.Member, address string) error {
	tx, err := db.MyDb.Begin()
	log.Println("init")
	queryMember := "INSERT INTO member (name, email) VALUES (?, ?)"
	result, err := tx.Exec(queryMember, member.Name, member.Email)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	memberID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	queryAuthority := "INSERT INTO authorities VALUES (?,?)"
	result, err = tx.Exec(queryAuthority, memberID, "ROLE_GUEST")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	queryRequest := "insert into member_request(member_id, address) VALUES (?,?)"
	_, err = tx.Exec(queryRequest, memberID, address)
	if err != nil {
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SaveAddress(id int, addr string) error {

	query := "insert into member_request(member_id, address) VALUES (?,?)"
	_, err := db.MyDb.Query(query, id, addr)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

	//tx, err := db.MyDb.Begin()
	//UpdateAddress := "UPDATE member SET address = ? WHERE id = ?"
	//_, err = tx.Exec(UpdateAddress, addr, id)
	//if err != nil {
	//	tx.Rollback()
	//	return "", fmt.Errorf("failed to update address: %v", err)
	//}
	//
	//token, err := jwt.AddressTokenProvider(id, addr)
	//if err != nil {
	//	tx.Rollback()
	//	return "", fmt.Errorf("failed to make addrToken: %v", err)
	//}
	//err = tx.Commit()
	//if err != nil {
	//	return token, err
	//}
	//return token, nil
}

func DeleteMember(id int) error {
	DeleteMemberQuery := "Delete from member where id = ?"
	_, err := db.MyDb.Exec(DeleteMemberQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func findAllMembers(page int) ([]models.MemberRequestOutput, error) {
	pageSize := 10
	offset := (page - 1) * pageSize
	query := `
		SELECT 
			mr.id AS request_id,
			m.name AS member_name,
			m.email AS member_email,
			m.id AS member_id,
			mr.address AS request_address,
			mr.state AS request_state,
			m.reg_date AS member_reg_date,
			mr.confirm_date AS request_confirm_date,
			a.authority AS authority_role
		FROM 
			member_request mr
		JOIN 
			member m ON mr.member_id = m.id
		LEFT JOIN 
			authorities a ON m.id = a.member_id
		limit ? offset ?
	`

	rows, err := db.MyDb.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	memberRequestsMap := make(map[int]*models.MemberRequestOutput)
	for rows.Next() {
		var (
			requestID          int
			memberName         string
			memberEmail        string
			memberId           int
			requestAddress     string
			requestState       int
			memberRegDate      time.Time
			requestConfirmDate time.Time
			authorityRole      sql.NullString
		)
		err := rows.Scan(
			&requestID,
			&memberName,
			&memberEmail,
			&memberId,
			&requestAddress,
			&requestState,
			&memberRegDate,
			&requestConfirmDate,
			&authorityRole,
		)
		if err != nil {
			return nil, err
		}

		if memberRequest, ok := memberRequestsMap[requestID]; ok {
			if authorityRole.Valid {
				memberRequest.Authorities = append(memberRequest.Authorities, models.Authority{Role: authorityRole.String})
			}
		} else {
			memberRequestsMap[requestID] = &models.MemberRequestOutput{
				RequestID:          requestID,
				MemberName:         memberName,
				MemberEmail:        memberEmail,
				MemberId:           memberId,
				RequestAddress:     requestAddress,
				RequestState:       requestState,
				MemberRegDate:      memberRegDate.Format("2006-01-02"),
				RequestConfirmDate: requestConfirmDate.Format("2006-01-02"),
				Authorities:        make([]models.Authority, 0),
			}
			if authorityRole.Valid {
				memberRequestsMap[requestID].Authorities = append(memberRequestsMap[requestID].Authorities, models.Authority{Role: authorityRole.String})
			}
		}
	}

	memberRequests := make([]models.MemberRequestOutput, 0, len(memberRequestsMap))
	for _, memberRequest := range memberRequestsMap {
		memberRequests = append(memberRequests, *memberRequest)
	}

	return memberRequests, nil
}

func FindMemberById(id int) (models.Member, error) {
	member := models.Member{}

	query := "SELECT * FROM member WHERE id=?"
	var addr sql.NullString
	err := db.MyDb.QueryRow(query, id).Scan(&member.ID, &member.Name, &member.Email, &addr, &member.RegDate)
	if addr.Valid {
		member.Address = addr.String
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return member, fmt.Errorf("member not found")
		}
		return member, err
	}
	return member, nil
}

func Confirm(id int, state int, memberId int, address string, authority string) error {

	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return err
	}

	tx, err := db.MyDb.Begin()
	if err != nil {
		return err
	}

	confirmDate := time.Now().In(loc)

	queryMemberRequest := "update member_request set state=?, confirm_date=? where id=?"
	_, err = tx.Exec(queryMemberRequest, state, confirmDate, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryAuthoritiesPrev := "delete from authorities where member_id=?"
	_, err = tx.Exec(queryAuthoritiesPrev, memberId)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryAuthorities := "insert into authorities(member_id, authority) VALUES (?,?)"
	_, err = tx.Exec(queryAuthorities, memberId, authority)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryMember := "update member set address=? where id=?"
	_, err = tx.Exec(queryMember, address, memberId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func findReqByMemberIdAndAddress(memberId int, address string) (models.MemberRequest, error) {
	result := models.MemberRequest{}
	query := "select * from member_request where member_id = ? and address = ?"
	var cDate sql.NullTime
	err := db.MyDb.QueryRow(query, memberId, address).Scan(&result.Id, &result.MemberId, &result.Address, &result.RegDate, &cDate, &result.State)
	if cDate.Valid {
		result.RegDate = cDate.Time
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return models.MemberRequest{}, fmt.Errorf("Req not found")
		}
		return models.MemberRequest{}, err
	}

	return result, err
}

func GetMemberCount() (int, error) {
	query := "select count(*) from member"
	var count int
	err := db.MyDb.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, err
}
