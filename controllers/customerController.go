package controllers

import (
	"final-project-ticketing-api/service"
	"final-project-ticketing-api/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	customer, err := service.GetCustomerById(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func CreateCustomer(c *gin.Context) {
	var request structs.CustomerRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service

	customer, errors := service.CreateCustomer(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert customer",
		"data":    customer,
	})
}

func Login(c *gin.Context) {
	var request structs.CustLogin
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service

	customer, errors := service.Login(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   customer.Token,
	})
}

func UpdateCustomer(c *gin.Context) {
	var request structs.CustomerRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	customer, errors := service.UpdateCustomer(request, id)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update customer with ID : " + strconv.Itoa(id),
		"data":    customer,
	})
}
