package internal

import "context"

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// GetAll returns all the tickets
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)

	// GetTicketsByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)
}
