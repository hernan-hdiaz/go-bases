package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id         int
	Name       string
	Email      string
	Country    string
	FlightTime int
	Price      int
}

var Tickets = []Ticket{}

func ConvertCSVtoTickets(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		var ticket Ticket
		if value, err := strconv.Atoi(line[0]); err == nil {
			ticket.Id = value
		}
		ticket.Name = line[1]
		ticket.Email = line[2]
		ticket.Country = line[3]
		if value, err := strconv.Atoi(strings.Split(line[4], ":")[0]); err == nil {
			ticket.FlightTime = value
		}
		if value, err := strconv.Atoi(line[5]); err == nil {
			ticket.Price = value
		}
		Tickets = append(Tickets, ticket)
	}
}

// Requerimiento 1

func GetTotalTickets(destination string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("not found data")
	}
	total := 0
	for _, v := range Tickets {
		if v.Country == destination {
			total++
		}
	}
	return total, nil
}

// Requerimiento 2

func GetCountByPeriod(time string) (int, error) {
	if len(Tickets) <= 0 {
		return 0, errors.New("not found data")
	}
	total := 0
	switch time {
	case "madrugada":
		for _, v := range Tickets {
			if int(v.FlightTime) >= 0 && v.FlightTime <= 6 {
				total++
			}
		}
		return total, nil
	case "mañana":
		for _, v := range Tickets {
			if int(v.FlightTime) >= 7 && v.FlightTime <= 12 {
				total++
			}
		}
		return total, nil
	case "tarde":
		for _, v := range Tickets {
			if int(v.FlightTime) >= 13 && v.FlightTime <= 19 {
				total++
			}
		}
		return total, nil
	case "noche":
		for _, v := range Tickets {
			if int(v.FlightTime) >= 20 && v.FlightTime <= 23 {
				total++
			}
		}
		return total, nil
	default:
		return 0, errors.New("not found period")
	}
}

// Requerimiento 3

func AverageDestination(destination string, total int) (float64, error) {
	totalDestinationTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total <= 0 {
		return 0, errors.New("can not divide by 0")
	}
	return float64(totalDestinationTickets) / float64(total), nil
}
