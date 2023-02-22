package service

import (
	"errors"
	"regexp"
	"ticketing/ticketing/database"
	"ticketing/ticketing/repository"
	"ticketing/ticketing/structs"
	"time"
)

func GetAllEventsByEventId(id int) ([]structs.TicketGet, error) {
	var result []structs.TicketGet
	err, result := repository.GetAllTicketByEventId(database.DBConnection, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetAllEvents() ([]structs.EventGet, error) {
	var result []structs.EventGet
	err, result := repository.GetAllEvent(database.DBConnection)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateEvent(request structs.EventRequest) (structs.Event, []error) {
	var err []error
	event, err := prepareRequestEvent(request)
	if err != nil {
		return event, err
	}
	event, err1 := repository.InsertEvent(database.DBConnection, event)
	if err1 != nil {
		err = append(err, err1)
		return event, err
	}
	return event, nil
}

func UpdateEvent(request structs.EventRequest, eventId int) (structs.Event, []error) {
	var result structs.Event
	var err []error
	err1, _ := repository.GetEventById(database.DBConnection, eventId)
	if err1 != nil {
		err = append(err, err1)
		return result, err
	}
	event, err := prepareRequestEvent(request)
	event.ID = eventId
	if err != nil {
		return event, err
	}
	event, err = repository.UpdateEvent(database.DBConnection, event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func DeleteEvent(eventId int) error {
	err, _ := repository.GetEventById(database.DBConnection, eventId)
	if err != nil {
		return err
	}
	err = repository.DeleteEvent(database.DBConnection, eventId)
	if err != nil {
		return err
	}
	return nil
}

func prepareRequestEvent(request structs.EventRequest) (structs.Event, []error) {
	var event structs.Event
	request, err, startTime, endTime := validateRequestEvent(request)
	if err != nil {
		return event, err
	}
	event.Name = request.Name
	event.Description = request.Description
	event.StartDate = startTime
	event.EndDate = endTime
	event.CategoryId = request.CategoryId
	return event, nil
}

func validateRequestEvent(request structs.EventRequest) (structs.EventRequest, []error, time.Time, time.Time) {
	var startInt []int
	var endInt []int
	var err []error
	startDate := request.StartDate
	endDate := request.StartDate
	regex, _ := regexp.Compile(formatDate)
	if !regex.MatchString(startDate) {
		err = append(err, errors.New("parameter 'start_date' must be in format yyyy-MM-dd"))
		panic(err)
	}
	if !regex.MatchString(endDate) {
		err = append(err, errors.New("parameter 'end_date' must be in format yyyy-MM-dd"))
		panic(err)
	}
	startTime, err1 := GetDate(startDate, startInt)
	if err1 != nil {
		err = append(err, errors.New("parameter 'start_date' must be in format yyyy-MM-dd"))
	}
	endTime, err2 := GetDate(startDate, endInt)
	if err2 != nil {
		err = append(err, errors.New("parameter 'end_date' must be in format yyyy-MM-dd"))
	}
	if !startTime.After(time.Now()) {
		err = append(err, errors.New("parameter 'start_date' cannot be yesterday"))
	}
	if endTime.Before(startTime) {
		err = append(err, errors.New("parameter 'end_date' must be greater than 'start_date'"))
	}
	if len(err) > 0 {
		return request, err, startTime, endTime
	}
	return request, nil, startTime, endTime
}
