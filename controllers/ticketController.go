package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetTicketById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	ticket, err := service2.GetTicketById(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}

func CreateTicket(c *gin.Context) {
	var request structs2.TicketRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service

	ticket, errors := service2.CreateTicket(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert ticket",
		"data":    ticket,
	})
}

func UpdateTicket(c *gin.Context) {
	var request structs2.TicketRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	ticket, errors := service2.UpdateTicket(request, id)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update ticket with ID : " + strconv.Itoa(id),
		"data":    ticket,
	})
}

func DeleteTicket(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service2.DeleteTicket(id)
	service2.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete ticket with ID : " + strconv.Itoa(id),
	})
}
