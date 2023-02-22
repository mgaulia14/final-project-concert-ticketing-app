package service

import (
	"errors"
	"regexp"
	"ticketing/ticketing/database"
	"ticketing/ticketing/repository"
	"ticketing/ticketing/structs"
	"time"
)

const formatDate = `\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])*`

func GetTicketById(ticketId int) (structs.Ticket, error) {
	var result structs.Ticket
	err, result := repository.GetByTicketId(database.DBConnection, ticketId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateTicket(request structs.TicketRequest) (structs.Ticket, []error) {
	ticket, err := prepareRequestTicket(request)
	if err != nil {
		return ticket, err
	}
	ticket, err = repository.InsertTicket(database.DBConnection, ticket)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

func UpdateTicket(request structs.TicketRequest, ticketId int) (structs.Ticket, []error) {
	var result structs.Ticket
	var err []error
	err1, _ := repository.GetByTicketId(database.DBConnection, ticketId)
	if err1 != nil {
		err = append(err, err1)
		return result, err
	}
	ticket, err := prepareRequestTicket(request)
	ticket.ID = ticketId
	if err != nil {
		return ticket, err
	}
	ticket, err = repository.UpdateTicket(database.DBConnection, ticket)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

func DeleteTicket(ticketId int) error {
	err, _ := repository.GetByTicketId(database.DBConnection, ticketId)
	if err != nil {
		return err
	}
	err = repository.DeleteTicket(database.DBConnection, ticketId)
	if err != nil {
		return err
	}
	return nil
}

func prepareRequestTicket(request structs.TicketRequest) (structs.Ticket, []error) {
	var ticket structs.Ticket
	request, err, dateTicket := validateRequestTicket(request)
	if err != nil {
		return ticket, err
	}
	ticket.Name = request.Name
	ticket.Date = dateTicket
	ticket.Price = request.Price
	ticket.Quota = request.Quota
	ticket.EventId = request.EventId
	return ticket, nil
}

func validateRequestTicket(request structs.TicketRequest) (structs.TicketRequest, []error, time.Time) {
	var dateInt []int
	var err []error
	var dateTicket time.Time
	dateRequest := request.Date

	// check is event id exist
	err1, ticket := repository.GetEventById(database.DBConnection, request.EventId)
	if err1 != nil {
		err = append(err, err1)
		return request, err, dateTicket
	}
	regex, _ := regexp.Compile(formatDate)
	if !regex.MatchString(dateRequest) {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
		panic(err)
	}
	dateTicket, err1 = GetDate(dateRequest, dateInt)
	if err1 != nil {
		err = append(err, errors.New("parameter 'date' must be in format yyyy-MM-dd"))
	}
	if !dateTicket.After(time.Now()) {
		err = append(err, errors.New("parameter 'date' cannot be yesterday"))
	}
	if dateTicket.Before(ticket.StartDate) || dateTicket.After(ticket.EndDate) {
		err = append(err, errors.New("parameter 'date' must between "+ticket.StartDate.String()+" or "+ticket.EndDate.String()))
	}
	if len(err) > 0 {
		return request, err, dateTicket
	}
	return request, nil, dateTicket
}
