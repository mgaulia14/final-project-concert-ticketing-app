package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	service2 "ticketing/service"
	structs2 "ticketing/structs"
)

func GetAllEvent(c *gin.Context) {
	// proses request to service
	events, err := service2.GetAllEvents()
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetAllTicketByEventId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// proses request to service
	tickets, err := service2.GetAllEventsByEventId(id)
	service2.CheckIsError(c, err)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func CreateEvent(c *gin.Context) {
	var request structs2.EventRequest
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service

	event, errors := service2.CreateEvent(request)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert event",
		"data":    event,
	})
}

func UpdateEvent(c *gin.Context) {
	var request structs2.EventRequest
	id, _ := strconv.Atoi(c.Param("id"))
	// bind JSON
	err := c.ShouldBindJSON(&request)
	service2.CheckIsError(c, err)

	// proses request to service
	event, errors := service2.UpdateEvent(request, id)
	service2.CheckIsErrors(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success update event with ID : " + strconv.Itoa(id),
		"data":    event,
	})
}

func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// proses request to service
	errors := service2.DeleteEvent(id)
	service2.CheckIsError(c, errors)

	// print success
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete event with ID : " + strconv.Itoa(id),
	})
}
