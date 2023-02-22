package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetAllCategory(c *gin.Context) {
	// proses request to service
	category, err := service2.GetAllCategory()
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func GetAllEventByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// proses request to service
	category, err := service2.GetAllEventsByCategory(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func CreateCategory(c *gin.Context) {
	var request structs2.CategoryRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	cat, err := service2.CreateCategory(request)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert category",
		"data":    cat,
	})
}

func UpdateCategory(c *gin.Context) {
	var request structs2.CategoryRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	cat, errs := service2.UpdateCategory(request, id)
	service2.CheckIsErrors(c, errs)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update category with ID : " + strconv.Itoa(id),
		"data":    cat,
	})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service2.DeleteCategory(id)
	service2.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete category with ID : " + strconv.Itoa(id),
	})
}
