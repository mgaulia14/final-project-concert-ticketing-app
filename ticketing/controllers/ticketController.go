package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticketing/ticketing/service"
	"ticketing/ticketing/structs"
)

func GetTicketById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	ticket, err := service.GetTicketById(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}

func CreateTicket(c *gin.Context) {
	var request structs.TicketRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service

	ticket, errors := service.CreateTicket(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert ticket",
		"data":    ticket,
	})
}

func UpdateTicket(c *gin.Context) {
	var request structs.TicketRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	ticket, errors := service.UpdateTicket(request, id)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update ticket with ID : " + strconv.Itoa(id),
		"data":    ticket,
	})
}

func DeleteTicket(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service.DeleteTicket(id)
	service.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete ticket with ID : " + strconv.Itoa(id),
	})
}
