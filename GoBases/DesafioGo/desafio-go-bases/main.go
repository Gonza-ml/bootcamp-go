package main

import (
	"desafio-go-bases/internal/tickets/service"
	"desafio-go-bases/internal/tickets/storage"
	"fmt"
	"strconv"
	"time"
)

// parse line to csv on a ticket
func parseTicket(data []string) (ticket storage.Ticket, err error) {
	// parse id
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return
	}

	// parse flight time
	flightTime, err := time.Parse("15:06", data[4])
	if err != nil {
		return
	}

	//parse price
	price, err := strconv.ParseFloat(data[5], 64)
	if err != nil {
		return
	}

	ticket = storage.Ticket{
		Id:                 id,
		Name:               data[1],
		Email:              data[2],
		DestinationCountry: data[3],
		FlightTime:         flightTime,
		Price:              price,
	}
	return
}

func main() {
	//init TicketStorage
	ticketMap := service.NewTicketMap(nil)

	// read tickets.csv
	dataTickets, err := storage.ReadCSV("tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	//save tickets
	for _, ticket := range dataTickets {
		parseTicket, err := parseTicket(ticket)
		if err != nil {
			fmt.Println("Hubo un error en el parseo.")
			fmt.Println(err)
			continue
		}
		err = ticketMap.SaveTicket(parseTicket)
		if err != nil {
			fmt.Println("Hubo un error en el guardado de ticket.")
			fmt.Println(err)
			fmt.Println("Ticket: ", parseTicket)
		}
	}

	// getTotalTickets
	destination := "Argentina"
	total, err := ticketMap.GetTotalTickets(destination)
	if err != nil {
		fmt.Println("Hubo un error en el total de tickets.")
		fmt.Println(err)
	}
	fmt.Println("Total de tickets con destino a", destination, ":", total)

	// GetCountByPeriod
	period := service.Night
	total, err = ticketMap.GetCountByPeriod(period)
	if err != nil {
		fmt.Println("Hubo un error en el total de tickets.")
		fmt.Println(err)
	}
	fmt.Println("Total de tickets en el horario", period, ":", total)

	// AverageDestination
	destination = "Argentina"
	avg, err := ticketMap.AverageDestination(destination)
	if err != nil {
		fmt.Println("Hubo un error en el porcentaje de personas.")
		fmt.Println(err)
	}
	fmt.Println("Porcentaje de personas que viajan a", destination, ": %", avg)
}
