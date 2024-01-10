package service_test

import (
	"desafio-go-bases/internal/tickets/service"
	"desafio-go-bases/internal/tickets/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("Success - case 01 - tickets with destination ", func(t *testing.T) {
		// arrange
		// ticketMapTest := storage.NewTicketMap(nil)
		// err := ticketMapTest.SaveTicket(tickets.Ticket{Id: 2, Name: "Jhon", DestinationCountry: "Finland"})
		// err = ticketMapTest.SaveTicket(tickets.Ticket{Id: 1, Name: "Jane", DestinationCountry: "Finland"})
		// err = ticketMapTest.SaveTicket(tickets.Ticket{Id: 3, Name: "Steve", DestinationCountry: "Brazil"})
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
		})

		// act
		destination := "Finland"
		total, err := ticketMapTest.GetTotalTickets(destination)

		// assert
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})

	t.Run("Success - case 02 - no ticket with destination ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
		})

		// act
		destination := "Argentina"
		total, err := ticketMapTest.GetTotalTickets(destination)
		// assert

		require.NoError(t, err)
		require.Equal(t, 0, total)
	})

	t.Run("Failure - case 03 - Invalid destination ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
		})

		// act
		destination := ""
		total, err := ticketMapTest.GetTotalTickets(destination)
		// assert
		expectedError := &service.ErrorTicketMap{
			Msg:     "Invalid destination",
			Feature: "destination",
		}
		expectedResult := 0

		require.Error(t, err)
		// require.ErrorIs(t, err, &service.ErrorTicketMap{})
		// require.EqualError(t, err, expectedError.Msg expectedError.Feature)
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedResult, total)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("Success - case 01 - tickets in to period ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", FlightTime: time.Date(0, 0, 0, 5, 59, 59, 0, time.UTC)},
			1: {Id: 2, Name: "Jane", FlightTime: time.Date(0, 0, 0, 6, 59, 50, 0, time.UTC)},
			2: {Id: 2, Name: "Steve", FlightTime: time.Date(0, 0, 0, 14, 50, 50, 0, time.UTC)},
		})

		// act
		period := service.Aftermoon
		total, err := ticketMapTest.GetCountByPeriod(period)
		// assert
		var expectedError error
		expectedResult := 2

		require.NoError(t, expectedError, err)
		require.Equal(t, expectedResult, total)
	})

	t.Run("Success - case 02 - no ticket in to period ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", FlightTime: time.Date(0, 0, 0, 5, 59, 59, 0, time.UTC)},
			1: {Id: 2, Name: "Jane", FlightTime: time.Date(0, 0, 0, 6, 59, 50, 0, time.UTC)},
			2: {Id: 2, Name: "Steve", FlightTime: time.Date(0, 0, 0, 14, 50, 50, 0, time.UTC)},
		})

		// act
		period := service.Night
		total, err := ticketMapTest.GetCountByPeriod(period)
		// assert
		var expectedError error
		expectedResult := 0

		require.NoError(t, expectedError, err)
		require.Equal(t, expectedResult, total)
	})

	t.Run("Failure - case 03 - Invalid period ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", FlightTime: time.Date(0, 0, 0, 5, 59, 59, 0, time.UTC)},
			1: {Id: 2, Name: "Jane", FlightTime: time.Date(0, 0, 0, 6, 59, 50, 0, time.UTC)},
			2: {Id: 2, Name: "Steve", FlightTime: time.Date(0, 0, 0, 14, 50, 50, 0, time.UTC)},
		})

		// act
		period := "invalido"
		total, err := ticketMapTest.GetCountByPeriod(period)
		// assert
		expectedError := &service.ErrorTicketMap{
			Msg:     "Invalid period",
			Feature: "time",
		}
		expectedResult := 0

		// require.Error(t, expectedError, err)
		require.Equal(t, expectedResult, total)
		require.Equal(t, expectedError, err)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Success - case 01 - people who travel to a destination ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
			4: {Id: 2, Name: "Ana", DestinationCountry: "Argentina"},
		})

		// act
		destination := "Finland"
		avg, err := ticketMapTest.AverageDestination(destination)
		// assert
		var expectedError error
		expectedResult := float64(50)

		require.NoError(t, expectedError, err)
		require.Equal(t, expectedResult, avg)
	})

	t.Run("Success - case 02 - any people travel a destination ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
			4: {Id: 2, Name: "Ana", DestinationCountry: "Argentina"},
		})

		// act
		destination := "Jamaica"
		avg, err := ticketMapTest.AverageDestination(destination)
		// assert
		var expectedError error
		expectedResult := float64(0)

		require.NoError(t, expectedError, err)
		require.Equal(t, expectedResult, avg)
	})

	t.Run("Failure - case 03 - Invalid destination ", func(t *testing.T) {
		// arrange
		ticketMapTest := service.NewTicketMap(map[int]storage.Ticket{
			0: {Id: 2, Name: "Jhon", DestinationCountry: "Finland"},
			1: {Id: 2, Name: "Jane", DestinationCountry: "Finland"},
			2: {Id: 2, Name: "Steve", DestinationCountry: "Brazil"},
			4: {Id: 2, Name: "Ana", DestinationCountry: "Argentina"},
		})

		// act
		destination := ""
		avg, err := ticketMapTest.AverageDestination(destination)
		// assert
		expectedError := &service.ErrorTicketMap{
			Msg:     "Invalid destination",
			Feature: "destination",
		}
		expectedResult := float64(0)

		// require.Error(t, expectedError, err)
		require.Equal(t, expectedResult, avg)
		require.Equal(t, expectedError, err)
	})
}
