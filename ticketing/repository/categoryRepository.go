package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"ticketing/ticketing/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sqlQuery := `SELECT * FROM category`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var category = structs.Category{}
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}
	return
}

func GetByCategoryById(db *sql.DB, catId int) (err error) {
	sqlQuery := "SELECT * FROM category WHERE id = $1"

	rows, _ := db.Query(sqlQuery, catId)

	if !rows.Next() {
		err = errors.New("category with ID : " + strconv.Itoa(catId) + " not found")
		return
	}
	return nil
}

func InsertCategory(db *sql.DB, category structs.Category) (structs.Category, error) {
	sqlQuery := `INSERT INTO category (name, created_at, updated_at) 
				VALUES ($1, $2, $3) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		category.Name,
		time.Now(),
		time.Now()).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)
	if err != nil {
		return category, err
	}
	return category, nil
}

func UpdateCategory(db *sql.DB, category structs.Category) (structs.Category, []error) {
	var errs []error
	sqlQuery := `UPDATE category 
				SET name = $1,
                    updated_at = $2
                WHERE id = $3`
	err := db.QueryRow(sqlQuery,
		category.Name,
		time.Now(),
		category.ID).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt)
	if errs != nil {
		errs = append(errs, err)
		return category, errs
	}
	return category, nil
}

func DeleteCategory(db *sql.DB, categoryId int) (err error) {
	sqlQuery := `DELETE FROM category WHERE id = $1`
	errs := db.QueryRow(sqlQuery, categoryId)
	return errs.Err()
}
