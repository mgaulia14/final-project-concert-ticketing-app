package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetWalletInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	wallet, err := service2.GetWalletInfo(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": wallet,
	})
}

func TopUpBalance(c *gin.Context) {
	var request structs2.WalletTopUp
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	wallet, errors := service2.TopUpWallet(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success top up wallet for account number: " + strconv.Itoa(request.AccountNumber),
		"data":    wallet,
	})
}
