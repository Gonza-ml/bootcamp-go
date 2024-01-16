package handler

import (
	"app/internal"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi"
)

// HandlerTicketDefault represents the default handler of the tickets
type HandlerTicketDefault struct {
	// srv represents the service of the tickets
	sv internal.ServiceTicket
}

func NewHandlerTicketDefault(sv internal.ServiceTicket) *HandlerTicketDefault {
	return &HandlerTicketDefault{
		sv: sv,
	}
}

func (h *HandlerTicketDefault) GetTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the country from the url
		country := chi.URLParam(r, "country")

		// check country
		if country == "" {
			response.Error(w, http.StatusBadRequest, "country is required")
		}

		// get the tickets by destination country
		totalTickets, err := h.sv.GetTicketsAmountByDestinationCountry(country)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrTicketGettingDest):
				response.Error(w, http.StatusInternalServerError, "Error getting the tickets")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
		}

		// send the response
		response.JSON(w, http.StatusOK, totalTickets)

	}
}

func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the country from the url
		country := chi.URLParam(r, "country")
		//check country
		if country == "" {
			response.Error(w, http.StatusBadRequest, "country is required")
		}

		// get the percentage of tickets by destination country
		avg, err := h.sv.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrTicketGettingDest):
				response.Error(w, http.StatusInternalServerError, "Error getting the tickets destination")
			case errors.Is(err, internal.ErrTicketGetting):
				response.Error(w, http.StatusInternalServerError, "Error getting the tickets")
			default:
				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
		}

		// send the response
		response.JSON(w, http.StatusOK, avg)
	}
}
