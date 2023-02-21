package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CheckIsError(c *gin.Context, err error) {
	if err != nil {
		// print error
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{
				"error_message": err.Error(),
				"error_status":  "Data not found",
			})
			panic(err)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
				"error_status":  "Invalid parameter",
			})
			panic(err)
		}
	}
}

func CheckIsErrors(c *gin.Context, err []error) {
	var errors []string
	if len(err) > 0 {
		// print error
		for _, e := range err {
			errors = append(errors, e.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": errors,
			"error_status":  "Invalid parameter",
		})
		panic(err)
	}
}

func GetDate(dateRequest string, dateInt []int) (time.Time, error) {
	var dateTicket time.Time
	dateString := strings.Split(dateRequest, "-")
	for _, i := range dateString {
		j, err := strconv.Atoi(i)
		if err != nil {
			return dateTicket, err
		}
		dateInt = append(dateInt, j)
	}
	dateTicket = GetDateTime(dateInt[0], dateInt[1], dateInt[2])
	return dateTicket, nil
}

func GetDateTime(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
