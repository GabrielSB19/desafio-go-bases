package tickets

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func setup() {
	// Limpiar el slice Tickets antes de cada prueba
	Tickets = []Ticket{}

	// Añadir algunos tickets de ejemplo
	AddTicket(1, "Tait Mc Caughan", "tmc0@scribd.com", "Finland", "17:11", 785)
	AddTicket(2, "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537)
	AddTicket(3, "Yalonda Jermyn", "yjermyn2@omniture.com", "China", "18:11", 579)
	AddTicket(4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238)
	AddTicket(5, "Jane Doe", "jdoe@example.com", "China", "05:00", 300)
}

func TestGetTotalTickets(t *testing.T) {
	setup()

	t.Run("Total Tickets for China", func(t *testing.T) {
		total, err := GetTotalTickets("China")
		require.NoError(t, err)
		require.Equal(t, 3, total)
	})

	t.Run("Total Tickets for Finland", func(t *testing.T) {
		total, err := GetTotalTickets("Finland")
		require.NoError(t, err)
		require.Equal(t, 1, total)
	})

	t.Run("Total Tickets for Mongolia", func(t *testing.T) {
		total, err := GetTotalTickets("Mongolia")
		require.NoError(t, err)
		require.Equal(t, 1, total)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	setup()

	t.Run("Counts by Period", func(t *testing.T) {
		madrugada, manana, tarde, noche, err := GetCountByPeriod()
		require.NoError(t, err)

		expectedCounts := []int{1, 0, 2, 2} // Madrugada: 1, Mañana: 0, Tarde: 2, Noche: 2
		actualCounts := []int{madrugada, manana, tarde, noche}

		require.Equal(t, expectedCounts, actualCounts)
	})

	t.Run("Counts with Different Data", func(t *testing.T) {
		// Se agregan nuevos datos para probar
		AddTicket(6, "Test User", "test@example.com", "Japan", "08:00", 100)
		madrugada, manana, tarde, noche, err := GetCountByPeriod()
		require.NoError(t, err)

		expectedCounts := []int{1, 1, 2, 2} // Actualizado: Madrugada: 1, Mañana: 1, Tarde: 2, Noche: 2
		actualCounts := []int{madrugada, manana, tarde, noche}

		require.Equal(t, expectedCounts, actualCounts)
	})

	t.Run("Empty Tickets", func(t *testing.T) {
		// Limpiar el slice de Tickets para probar un caso vacío
		Tickets = []Ticket{}
		madrugada, manana, tarde, noche, err := GetCountByPeriod()
		require.NoError(t, err)

		expectedCounts := []int{0, 0, 0, 0} // Todos los periodos deben ser 0
		actualCounts := []int{madrugada, manana, tarde, noche}

		require.Equal(t, expectedCounts, actualCounts)
	})
}

func TestAverageDestination(t *testing.T) {
	setup()

	t.Run("Average for China", func(t *testing.T) {
		average, err := AverageDestination("China")
		require.NoError(t, err)

		expectedAverage := float64(3) / float64(len(Tickets))
		require.Equal(t, expectedAverage, average)
	})

	t.Run("Average for Finland", func(t *testing.T) {
		average, err := AverageDestination("Finland")
		require.NoError(t, err)

		expectedAverage := float64(1) / float64(len(Tickets))
		require.Equal(t, expectedAverage, average)
	})

	t.Run("Average for Mongolia", func(t *testing.T) {
		average, err := AverageDestination("Mongolia")
		require.NoError(t, err)

		expectedAverage := float64(1) / float64(len(Tickets))
		require.Equal(t, expectedAverage, average)
	})
}