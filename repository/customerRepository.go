package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"ticketing/ticketing/structs"
	"time"
)

func GetByCustomerId(db *sql.DB, customerId int) (err error, result structs.Customer) {
	sqlQuery := `SELECT * FROM customer
				WHERE customer.id = $1`
	var customer = structs.Customer{}
	rows, err := db.Query(sqlQuery, customerId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&customer.ID,
			&customer.FullName,
			&customer.BirthDate,
			&customer.Address,
			&customer.Email,
			&customer.PhoneNumber,
			&customer.Password,
			&customer.CreatedAt,
			&customer.UpdatedAt,
			&customer.IsAdmin)
		if err != nil {
			panic(err)
		}
		result = customer
		return nil, customer
	}
	err = errors.New("customer with ID : " + strconv.Itoa(customerId) + " not found")
	return err, customer
}

func GetCustomerByEmail(db *sql.DB, email string) (err error, result structs.Customer) {
	sqlQuery := `SELECT * FROM customer
				WHERE customer.email = $1`
	var customer = structs.Customer{}
	rows, err := db.Query(sqlQuery, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&customer.ID,
			&customer.FullName,
			&customer.BirthDate,
			&customer.Address,
			&customer.Email,
			&customer.PhoneNumber,
			&customer.Password,
			&customer.CreatedAt,
			&customer.UpdatedAt,
			&customer.IsAdmin)
		if err != nil {
			panic(err)
		}
		result = customer
		return nil, customer
	}
	err = errors.New("customer with email : " + email + " not found")
	return err, customer
}

func InsertCustomer(db *sql.DB, customer structs.Customer) (structs.Customer, []error) {
	var errs []error
	sqlQuery := `INSERT INTO customer (
                      full_name, 
                      birth_date,
                      address,
                      phone_number,
                      email,
                      password,
                      created_at, 
                      updated_at,
                      isadmin) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		customer.FullName,
		customer.BirthDate,
		customer.Address,
		customer.PhoneNumber,
		customer.Email,
		customer.Password,
		time.Now(),
		time.Now(),
		customer.IsAdmin).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.BirthDate,
		&customer.Address,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Password,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.IsAdmin,
	)
	if err != nil {
		errs = append(errs, err)
		return customer, errs
	}
	return customer, nil
}

func UpdateCustomer(db *sql.DB, customer structs.Customer) (structs.Customer, []error) {
	var errs []error
	sqlQuery := `UPDATE customer 
				SET full_name = $1, 
				    birth_date = $2,
				    address = $3, 
				    phone_number = $4,
				    email = $5,
				    password = $6,
				    updated_at = $7
				WHERE id = $8`

	err := db.QueryRow(sqlQuery,
		customer.FullName,
		customer.BirthDate,
		customer.Address,
		customer.PhoneNumber,
		customer.Email,
		customer.Password,
		time.Now(),
		time.Now()).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.BirthDate,
		&customer.Address,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Password,
		&customer.CreatedAt,
		&customer.UpdatedAt,
		&customer.IsAdmin)
	if errs != nil {
		errs = append(errs, err)
		return customer, errs
	}
	return customer, nil
}
