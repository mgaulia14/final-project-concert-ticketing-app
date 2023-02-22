package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetTransactionById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	transaction, err := service2.GetTransactionById(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func GetTransactionByCustomerId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	transaction, err := service2.GetTransactionByCustomerId(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func GetAllTransactions(c *gin.Context) {
	// proses request to service
	transaction, err := service2.GetAllTransaction()
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func CreateTransaction(c *gin.Context) {
	var request structs2.TransactionRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	transaction, errors := service2.CreateTransaction(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert transaction",
		"data":    transaction,
	})
}
