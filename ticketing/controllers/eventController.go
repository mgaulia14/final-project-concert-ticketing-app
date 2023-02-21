package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ticketing/ticketing/service"
	"ticketing/ticketing/structs"
)

func GetAllEvent(c *gin.Context) {
	// proses request to service
	events, err := service.GetAllEvents()
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetAllTicketByEventId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// proses request to service
	tickets, err := service.GetAllEventsByEventId(id)
	service.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func CreateEvent(c *gin.Context) {
	var request structs.EventRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service

	event, errors := service.CreateEvent(request)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert event",
		"data":    event,
	})
}

func UpdateEvent(c *gin.Context) {
	var request structs.EventRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service.CheckIsError(c, err)

	// proses request to service
	event, errors := service.UpdateEvent(request, id)
	service.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update event with ID : " + strconv.Itoa(id),
		"data":    event,
	})
}

func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service.DeleteEvent(id)
	service.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete event with ID : " + strconv.Itoa(id),
	})
}
