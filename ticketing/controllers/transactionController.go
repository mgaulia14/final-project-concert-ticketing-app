package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticketing/ticketing/service"
	"ticketing/ticketing/structs"
)

func GetTransactionById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	transaction, err := service.GetTransactionById(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func GetTransactionByCustomerId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	transaction, err := service.GetTransactionByCustomerId(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func GetAllTransactions(c *gin.Context) {
	// proses request to service
	transaction, err := service.GetAllTransaction()
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": transaction,
	})
}

func CreateTransaction(c *gin.Context) {
	var request structs.TransactionRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	transaction, errors := service.CreateTransaction(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert transaction",
		"data":    transaction,
	})
}
