package service

import (
	"app/internal"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalAmountTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	ctx := context.Background()
	tickets, err := s.rp.Get(ctx)
	if err != nil {
		err = internal.ErrTicketGetting
		return
	}
	total = len(tickets)
	return
}

// GetTotalAmountTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	ctx := context.Background()
	tickets, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		err = internal.ErrTicketGettingDest
		return
	}
	total = len(tickets)
	return
}

// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(country string) (total int, err error) {
	ctx := context.Background()
	// Get tickets by destination country
	ticketsDest, err := s.rp.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		err = internal.ErrTicketGettingDest
		return
	}
	// Get all tickets
	ticketsTotal, err := s.rp.Get(ctx)
	if err != nil {
		err = internal.ErrTicketGetting
		return
	}
	total = int(float64(len(ticketsDest)) / float64(len(ticketsTotal)) * 100)
	return
}
