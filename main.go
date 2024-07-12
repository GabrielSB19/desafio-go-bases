package main

import (
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"fmt"
)

func main() {
	tickets.ReadFile("./tickets.csv")
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("Total tickets to Brazil:", total)
	madrugada, manana, tarde, noche, _ := tickets.GetCountByPeriod()
	fmt.Println("Tickets by period:")
	fmt.Println("Madrugada:", madrugada)
	fmt.Println("Mañana:", manana)
	fmt.Println("Tarde:", tarde)
	fmt.Println("Noche:", noche)

	average, _ := tickets.AverageDestination("Colombia")
	fmt.Println("Average tickets to Colombia:", average)
}