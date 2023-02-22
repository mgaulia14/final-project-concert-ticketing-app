package service

import (
	"final-project-ticketing-api/database"
	"final-project-ticketing-api/dto"
	"final-project-ticketing-api/repository"
	"final-project-ticketing-api/structs"
)

func GetAllCategory() ([]structs.Category, error) {
	var result []structs.Category
	err, result := repository.GetAllCategory(database.DBConnection)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllEventsByCategory(id int) ([]dto.EventGet, error) {
	var result []dto.EventGet
	err, result := repository.GetAllEventByCategoryId(database.DBConnection, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateCategory(request structs.CategoryRequest) (structs.Category, error) {
	cat := prepareRequestCategory(request)
	cat, err := repository.InsertCategory(database.DBConnection, cat)
	if err != nil {
		return cat, err
	}
	return cat, nil
}

func UpdateCategory(request structs.CategoryRequest, categoryId int) (structs.Category, []error) {
	var result structs.Category
	var err []error
	err1 := repository.GetByCategoryById(database.DBConnection, categoryId)
	if err1 != nil {
		err = append(err, err1)
		return result, err
	}
	category := prepareRequestCategory(request)
	category.ID = categoryId
	if err != nil {
		return category, err
	}
	category, err = repository.UpdateCategory(database.DBConnection, category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func DeleteCategory(categoryId int) error {
	err := repository.GetByCategoryById(database.DBConnection, categoryId)
	if err != nil {
		return err
	}
	err = repository.DeleteCategory(database.DBConnection, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func prepareRequestCategory(request structs.CategoryRequest) structs.Category {
	var cat structs.Category
	cat.Name = request.Name
	return cat
}
