package tickets

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Ticket struct {
	ID 	int
	Name string
	Email string
	Destination string
	FlightTime string
	Price float64
}

var Tickets = []Ticket{}

func ReadFile(path string) []Ticket {
	defer func(){
		if err:=recover(); err!=nil{
			fmt.Println("Error:", err)
		}
	}()

	file, err := os.Open(path)

	if err != nil{
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		item := scanner.Text()
		parts := strings.Split(item, ",")
		id, errI := strconv.Atoi(parts[0])
		price, errP := strconv.ParseFloat(parts[5], 64)

		if errI != nil || errP != nil {
			handlePanic("Error to convert the number")
		}

		AddTicket(id, parts[1], parts[2], parts[3], parts[4], price)
	}

	return Tickets
}

func AddTicket(id int, name string, email string, destination string, flightTime string, price float64) {
	Tickets = append(Tickets, Ticket{id, name, email, destination, flightTime, price})
}

func GetTotalTickets(destination string) (int, error) {
	total := 0
	for _, ticket := range Tickets {
		if ticket.Destination == destination {
			total++
		}
	}
	return total, nil
}

func GetCountByPeriod() (madrugada, manana, tarde, noche int, e error) {
	for _, ticket := range Tickets {
		hour, err := getHour(ticket.FlightTime)
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid time format")
		}

		switch {
			case hour >= 0 && hour <= 6:
				madrugada++
			case hour >= 7 && hour <= 12:
				manana++
			case hour >= 13 && hour <= 19:
				tarde++
			case hour >= 20 && hour <= 23:
				noche++
		}
	}
	return madrugada, manana, tarde, noche, nil
}

func getHour(time string)(int, error){
	parts := strings.Split(time, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid time format")
	}
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		handlePanic("Error to convert the number")
	}
	return hour, nil
}

func AverageDestination(destination string) (float64, error) {
	totalDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0, fmt.Errorf("error to get total tickets")
	
	}
	average := (float64(totalDestination) / float64(len(Tickets)))*100
	return average, nil
}

func handlePanic(msg string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic(msg)
}
