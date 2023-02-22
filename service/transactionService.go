package service

import (
	"errors"
	"final-project-ticketing-api/database"
	"final-project-ticketing-api/dto"
	"final-project-ticketing-api/repository"
	"final-project-ticketing-api/structs"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func GetTransactionById(transactionId int) (dto.TransactionGet, error) {
	var result dto.TransactionGet
	err, result := repository.GetByTransactionId(database.DBConnection, transactionId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetTransactionByCustomerId(customerId int) ([]dto.TransactionGet, error) {
	var result []dto.TransactionGet
	err, result := repository.GetTransactionsByCustomerId(database.DBConnection, customerId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllTransaction() ([]dto.TransactionGet, error) {
	var result []dto.TransactionGet
	err, result := repository.GetAllTransaction(database.DBConnection)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateTransaction(request structs.TransactionRequest) (structs.Transaction, []error) {
	transaction, err, ticket := prepareRequestTransaction(request)
	if err != nil {
		return transaction, err
	}
	transaction, err = repository.InsertTransaction(database.DBConnection, transaction)
	if err != nil {
		return transaction, err
	}
	ticket.Quota = ticket.Quota - 1
	ticket, errs := repository.UpdateTicket(database.DBConnection, ticket)
	if errs != nil {
		return transaction, err
	}
	return transaction, nil
}

func prepareRequestTransaction(request structs.TransactionRequest) (structs.Transaction, []error, structs.Ticket) {
	var transaction structs.Transaction
	request, err, dateTransaction, ticket, cust := validateRequestTransaction(request)
	if err != nil {
		return transaction, err, ticket
	}
	// generate qr code image
	qrCode := GenerateUniqueCode(6)
	err1 := qrcode.WriteFile("SUCCESS_"+"_"+ticket.Name+"_"+cust.FullName+"_"+qrCode, qrcode.Medium, 256, qrCode+".png")
	if err1 != nil {
		err = append(err, err1)
		return transaction, err, ticket
	}
	// access file local
	image, err1 := os.Open(qrCode + ".png")
	if err1 != nil {
		err = append(err, err1)
		return transaction, err, ticket
	}
	// upload image to CDN
	uploadUrl, err2 := ImageUploadHelper(image)
	if err2 != nil {
		err = append(err, err2)
		return transaction, err, ticket

	}
	transaction.Date = dateTransaction
	transaction.QrCode = uploadUrl
	transaction.TicketId = request.TicketId
	transaction.CustomerId = request.CustomerId
	return transaction, nil, ticket
}

func validateRequestTransaction(request structs.TransactionRequest) (structs.TransactionRequest, []error, time.Time, structs.Ticket, structs.Customer) {
	var dateInt []int
	var err []error
	var dateTransaction time.Time
	var cust structs.Customer
	dateRequest := request.Date
	err1, ticket := repository.GetByTicketId(database.DBConnection, request.TicketId)
	if err1 != nil {
		err = append(err, err1)
		return request, err, dateTransaction, ticket, cust
	}
	err1, cust = repository.GetByCustomerId(database.DBConnection, request.CustomerId)
	if err1 != nil {
		err = append(err, err1)
		return request, err, dateTransaction, ticket, cust
	}
	regex, _ := regexp.Compile(formatDate)
	if !regex.MatchString(dateRequest) {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
		panic(err)
	}
	dateTransaction, err1 = GetDate(dateRequest, dateInt)
	if err1 != nil {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
	}
	if !dateTransaction.After(time.Now()) {
		err = append(err, errors.New("parameter 'date' cannot less than today"))
	}
	err1, wallet := repository.GetWalletByCustomerId(database.DBConnection, cust.ID)
	if err1 != nil {
		err = append(err, err1)
	}
	ticketPrice, err1 := strconv.ParseFloat(ticket.Price, 8)
	if err1 != nil {
		err = append(err, err1)
	}
	if ticketPrice > wallet.Balance {
		err = append(err, errors.New("insufficient balance"))
	}
	if ticket.Quota == 0 {
		err = append(err, errors.New("ticket is sold out"))
	}

	if len(err) > 0 {
		return request, err, dateTransaction, ticket, cust
	}
	return request, nil, dateTransaction, ticket, cust
}

func GenerateUniqueCode(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
