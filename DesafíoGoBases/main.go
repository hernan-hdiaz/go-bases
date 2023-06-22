package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	tickets.ConvertCSVtoTickets("tickets.csv")

	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	println("Tickets con destino a Brasil: ", total)

	total, err = tickets.GetCountByPeriod("madrugada")
	if err != nil {
		fmt.Println(err)
	}
	println("Pasajeros de madrugada: ", total)

	total, _ = tickets.GetCountByPeriod("mañana")
	println("Pasajeros de mañana: ", total)
	if err != nil {
		fmt.Println(err)
	}

	total, err = tickets.GetCountByPeriod("tarde")
	println("Pasajeros de tarde: ", total)
	if err != nil {
		fmt.Println(err)
	}

	total, err = tickets.GetCountByPeriod("noche")
	println("Pasajeros de noche: ", total)
	if err != nil {
		fmt.Println(err)
	}

	totales, err := tickets.AverageDestination("Brazil", len(tickets.Tickets))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Promedio de pasajeros que viajan a Brasil en un día: ", totales)

}
