package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	customer, err := service2.GetCustomerById(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func CreateCustomer(c *gin.Context) {
	var request structs2.CustomerRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service

	customer, errors := service2.CreateCustomer(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert customer",
		"data":    customer,
	})
}

func Login(c *gin.Context) {
	var request structs2.CustLogin
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service

	customer, errors := service2.Login(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token":   customer.Token,
	})
}

func UpdateCustomer(c *gin.Context) {
	var request structs2.CustomerRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	customer, errors := service2.UpdateCustomer(request, id)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update customer with ID : " + strconv.Itoa(id),
		"data":    customer,
	})
}
