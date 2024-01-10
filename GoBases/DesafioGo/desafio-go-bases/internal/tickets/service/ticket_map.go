package service

import (
	"desafio-go-bases/internal/tickets/storage"
	"fmt"
	"time"
)

type ErrorTicketMap struct {
	Msg     string
	Feature string
}

func (e *ErrorTicketMap) Error() string {
	return fmt.Sprintf("%s. field: %s", e.Msg, e.Feature)
}

type Period string

const (
	Aftermoon = "madrugada"
	Morning   = "mañana"
	Evening   = "tarde"
	Night     = "noche"
)

type TicketsMap struct {
	tickets map[int]storage.Ticket
}

func NewTicketMap(ticketsMap map[int]storage.Ticket) *TicketsMap {
	defaultMap := make(map[int]storage.Ticket)

	if ticketsMap != nil {
		defaultMap = ticketsMap
	}

	return &TicketsMap{
		tickets: defaultMap,
	}
}

// Save ticket into TicketMap
func (t *TicketsMap) SaveTicket(ticket storage.Ticket) (err error) {
	// Check ticket
	if ticket.Id == 0 {
		err = &ErrorTicketMap{
			Msg:     "Invalid Ticket",
			Feature: "ID",
		}
		return
	}

	if ticket.Name == "" {
		err = &ErrorTicketMap{
			Msg:     "Invalid Ticket",
			Feature: "Name",
		}
		return
	}

	// OTHER CHEKS FIELD
	// ..

	t.tickets[ticket.Id] = ticket
	return
}

// GetTotalTickets get the total tickets with an specific destination
func (t *TicketsMap) GetTotalTickets(destination string) (total int, err error) {
	// check destination
	if destination == "" {
		err = &ErrorTicketMap{
			Msg:     "Invalid destination",
			Feature: "destination",
		}
		return
	}

	for _, ticket := range t.tickets {
		if ticket.DestinationCountry == destination {
			total++
		}
	}

	return
}

// GetCountByPeriod get the total number of people travel in a period
func (t *TicketsMap) GetCountByPeriod(period string) (total int, err error) {

	// initialization period
	var initPeriod time.Time
	var endPeriod time.Time

	switch period {
	case Aftermoon:
		initPeriod = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)  // 00:00
		endPeriod = time.Date(0, 0, 0, 6, 59, 59, 0, time.UTC) // 06:59
	case Morning:
		initPeriod = time.Date(0, 0, 0, 7, 0, 0, 0, time.UTC)   // 07:00
		endPeriod = time.Date(0, 0, 0, 12, 59, 59, 0, time.UTC) // 12:59
	case Evening:
		initPeriod = time.Date(0, 0, 0, 13, 0, 0, 0, time.UTC)  // 13:00
		endPeriod = time.Date(0, 0, 0, 19, 59, 59, 0, time.UTC) // 19:59
	case Night:
		initPeriod = time.Date(0, 0, 0, 20, 0, 0, 0, time.UTC)  // 20:00
		endPeriod = time.Date(0, 0, 0, 23, 59, 59, 0, time.UTC) // 23:59
	default:
		err = &ErrorTicketMap{
			Msg:     "Invalid period",
			Feature: "time",
		}
		return
	}
	// count if fligh time is in the range
	for _, ticket := range t.tickets {
		// create equivalent date
		flightTime := time.Date(0, 0, 0, ticket.FlightTime.Hour(), ticket.FlightTime.Minute(), 0, 0, time.UTC)
		if flightTime.After(initPeriod) && flightTime.Before(endPeriod) {
			total++
		}
	}
	return
}

// AverageDestination get the average to people travel a destination
func (t *TicketsMap) AverageDestination(destination string) (avg float64, err error) {
	// check destination
	if destination == "" {
		err = &ErrorTicketMap{
			Msg:     "Invalid destination",
			Feature: "destination",
		}
		return
	}
	// init count
	total := 0.0
	// compare destinations
	for _, ticket := range t.tickets {
		if ticket.DestinationCountry == destination {
			total++
		}
	}
	// get average
	avg = total / float64(len(t.tickets)) * 100
	return
}
