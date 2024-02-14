package transaction

import (
	"backend/internal/db"
	"backend/internal/models"
	"fmt"
)

func GetTransactionCount() (int, error) {
	query := "SELECT COUNT(*) FROM transaction"
	var count int
	err := db.MyDb.QueryRow(query).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
func DeleteTransaction(transactionId int) error {
	query := "delete from transaction where id=?"
	_, err := db.MyDb.Exec(query, transactionId)
	if err != nil {
		return err
	}
	return nil
}

func PutTransaction(transaction *models.Transaction) error {
	query := "UPDATE transaction SET title=?, amount=?, client=?, sell_buy=? WHERE id=?"
	_, err := db.MyDb.Exec(query, transaction.Title, transaction.Amount, transaction.Client, transaction.SellBuy, transaction.ID)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}
	return nil
}

func PostTransaction(transaction *models.Transaction) error {
	query := "INSERT INTO transaction (title, amount, client, sell_buy) VALUES (?, ?, ?, ?)"
	_, err := db.MyDb.Exec(query, transaction.Title, transaction.Amount, transaction.Client, transaction.SellBuy)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}
	return nil
}

func GetTransactions(page int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	limit := 10                  // 한 페이지당 표시할 트랜잭션 수
	offset := (page - 1) * limit // 오프셋 계산

	query := "SELECT id, title, amount, client, sell_buy, reg_date FROM transaction LIMIT ? OFFSET ?"
	rows, err := db.MyDb.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Title, &transaction.Amount, &transaction.Client, &transaction.SellBuy, &transaction.RegDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over transaction rows: %v", err)
	}

	return transactions, nil
}
