package internal

import "errors"

var (
	// ErrTicketNotFound represents the error when the ticket is not found
	ErrTicketGetting = errors.New("Error getting the tickets")
	// ErrTicketGettingDest represents the error when the ticket is not found with the destination country
	ErrTicketGettingDest = errors.New("Error getting the tickets by destination country")
)

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets() (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(country string) (total int, err error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(country string) (total int, err error)
}
