package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticketing/ticketing/service"
	"ticketing/ticketing/structs"
)

func GetAllCategory(c *gin.Context) {
	// proses request to service
	category, err := service.GetAllCategory()
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func GetAllEventByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// proses request to service
	category, err := service.GetAllEventsByCategory(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	var request structs.CategoryRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	cat, err := service.CreateCategory(request)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert category",
		"data":    cat,
	})
}

func UpdateCategory(c *gin.Context) {
	var request structs.CategoryRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	cat, errs := service.UpdateCategory(request, id)
	service.CheckIsErrors(c, errs)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update category with ID : " + strconv.Itoa(id),
		"data":    cat,
	})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service.DeleteCategory(id)
	service.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete category with ID : " + strconv.Itoa(id),
	})
}
