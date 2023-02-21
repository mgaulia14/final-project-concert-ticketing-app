package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticketing/ticketing/service"
	"ticketing/ticketing/structs"
)

func GetWalletInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	wallet, err := service.GetWalletInfo(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": wallet,
	})
}

func TopUpBalance(c *gin.Context) {
	var request structs.WalletTopUp
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	wallet, errors := service.TopUpWallet(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success top up wallet for account number: " + strconv.Itoa(request.AccountNumber),
		"data":    wallet,
	})
}
