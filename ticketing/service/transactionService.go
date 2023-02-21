package service

import (
	"errors"
	"regexp"
	"strconv"
	"ticketing/ticketing/database"
	"ticketing/ticketing/repository"
	"ticketing/ticketing/structs"
	"time"
)

func GetTransactionById(transactionId int) (structs.TransactionGet, error) {
	var result structs.TransactionGet
	err, result := repository.GetByTransactionId(database.DBConnection, transactionId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateTransaction(request structs.TransactionRequest) (structs.Transaction, []error) {
	transaction, err := prepareRequestTransaction(request)
	if err != nil {
		return transaction, err
	}
	transaction, err = repository.InsertTransaction(database.DBConnection, transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func prepareRequestTransaction(request structs.TransactionRequest) (structs.Transaction, []error) {
	var transaction structs.Transaction
	request, err, dateTransaction, ticket, cust := validateRequestTransaction(request)
	if err != nil {
		return transaction, err
	}
	transaction.Date = dateTransaction
	transaction.QrCode = "qrcode" // generate qr code
	transaction.TicketId = ticket
	transaction.CustomerId = cust
	return transaction, nil
}

func validateRequestTransaction(request structs.TransactionRequest) (structs.TransactionRequest, []error, time.Time, structs.Ticket, structs.Customer) {
	var dateInt []int
	var err []error
	dateRequest := request.Date
	regex, _ := regexp.Compile(formatDate)
	if !regex.MatchString(dateRequest) {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
		panic(err)
	}
	dateTransaction, err1 := GetDate(dateRequest, dateInt)
	if err1 != nil {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
	}
	if !dateTransaction.After(time.Now()) {
		err = append(err, errors.New("parameter 'date' cannot be yesterday"))
	}
	err1, ticket := repository.GetByTicketId(database.DBConnection, request.TicketId)
	if err1 != nil {
		err = append(err, err1)
	}
	err1, cust := repository.GetByCustomerId(database.DBConnection, request.CustomerId)
	if err1 != nil {
		err = append(err, err1)
	}
	err1, wallet := repository.GetWalletInfo(database.DBConnection, cust.WalletId)
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
