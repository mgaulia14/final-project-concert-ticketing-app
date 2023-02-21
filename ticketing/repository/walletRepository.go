package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"ticketing/ticketing/structs"
	"time"
)

func GetWalletInfo(db *sql.DB, walletId int) (err error, result structs.Wallet) {
	sqlQuery := `SELECT * FROM wallet
				WHERE wallet.id = $1`
	var wallet = structs.Wallet{}
	rows, err := db.Query(sqlQuery, walletId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&wallet.ID,
			&wallet.Balance,
			&wallet.AccountName,
			&wallet.AccountNumber,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		result = wallet
		return nil, wallet
	}
	err = errors.New("wallet with ID : " + strconv.Itoa(walletId) + " not found")
	return err, wallet
}

func GetWalletByAccountNumber(db *sql.DB, accountNumber int) (err error, result structs.Wallet) {
	sqlQuery := `SELECT * FROM wallet
				WHERE wallet.account_number = $1`
	var wallet = structs.Wallet{}
	rows, err := db.Query(sqlQuery, accountNumber)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&wallet.ID,
			&wallet.Balance,
			&wallet.AccountName,
			&wallet.AccountNumber,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		result = wallet
		return nil, wallet
	}
	err = errors.New("wallet with account number : " + strconv.Itoa(accountNumber) + " not found")
	return err, wallet
}

func InsertWallet(db *sql.DB, wallet structs.Wallet) (structs.Wallet, []error) {
	var errs []error
	sqlQuery := `INSERT INTO wallet (
                    balance, 
                    account_name, 
                    account_number,
                    created_at, 
                    updated_at) 
				VALUES ($1, $2, $3, $4, $5) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		wallet.Balance,
		wallet.AccountName,
		wallet.AccountNumber,
		time.Now(),
		time.Now()).Scan(
		&wallet.ID,
		&wallet.Balance,
		&wallet.AccountName,
		&wallet.AccountNumber,
		&wallet.CreatedAt,
		&wallet.UpdatedAt)
	if err != nil {
		errs = append(errs, err)
		return wallet, errs
	}
	return wallet, nil
}

func TopUpBalance(db *sql.DB, wallet structs.Wallet) (structs.Wallet, []error) {
	var errs []error
	sqlQuery := `UPDATE wallet 
				SET balance = $1,
				    updated_at = $2
				WHERE account_number = $3`

	err := db.QueryRow(sqlQuery,
		wallet.Balance,
		time.Now(),
		wallet.AccountNumber).Scan(
		&wallet.ID,
		&wallet.Balance,
		&wallet.AccountName,
		&wallet.AccountNumber,
		&wallet.CreatedAt,
		&wallet.UpdatedAt)
	if errs != nil {
		errs = append(errs, err)
		return wallet, errs
	}
	return wallet, nil
}

func DeleteWallet(db *sql.DB, id int) (err error) {
	sqlQuery := `DELETE FROM wallet WHERE id = $1`
	errs := db.QueryRow(sqlQuery, id)
	return errs.Err()
}
