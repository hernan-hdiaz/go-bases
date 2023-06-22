package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	ConvertCSVtoTickets("../../tickets.csv")
	t.Run("should get total persons for a country", func(t *testing.T) {
		//Arrange
		country := "Brazil"
		expect := 45

		//Act
		totalResult, err := GetTotalTickets(country)

		//Assert
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
	t.Run("should get error for country", func(t *testing.T) {
		//Arrange
		country := "Guya"
		expect := 0

		//Act
		totalResult, err := GetTotalTickets(country)

		//Assert
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	ConvertCSVtoTickets("../../tickets.csv")
	t.Run("should get total persons for a period", func(t *testing.T) {
		//Arrange
		time := "mañana"
		expect := 256

		//Act
		totalResult, err := GetCountByPeriod(time)

		//Assert
		assert.Equal(t, expect, totalResult)
		assert.Nil(t, err)
	})
	t.Run("should get total 0 for period medianoche", func(t *testing.T) {
		//Arrange
		time := "medianoche"
		expect := 0

		//Act
		totalResult, err := GetCountByPeriod(time)

		//Assert
		assert.Equal(t, expect, totalResult)
		assert.NotNil(t, err)
	})
}

func TestAverageDestination(t *testing.T) {
	ConvertCSVtoTickets("../../tickets.csv")
	t.Run("should get average destination for Brazil", func(t *testing.T) {
		//Arrange
		country := "Brazil"
		total := 10
		expect := 4.5

		//Act
		averageResult, err := AverageDestination(country, total)

		//Assert
		assert.Equal(t, expect, averageResult)
		assert.Nil(t, err)
	})
	t.Run("should get average 0 for destination Guya", func(t *testing.T) {
		//Arrange
		country := "Guya"
		total := 10
		expect := 0.0

		//Act
		averageResult, err := AverageDestination(country, total)

		//Assert
		assert.Equal(t, expect, averageResult)
		assert.Nil(t, err)
	})
	t.Run("should get error can't divide by zero ", func(t *testing.T) {
		//Arrange
		country := "Guya"
		total := 0
		expect := 0.0

		//Act
		averageResult, err := AverageDestination(country, total)

		//Assert
		assert.Equal(t, expect, averageResult)
		assert.NotNil(t, err)
	})
}
