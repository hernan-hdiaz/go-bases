package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 3 - Test del salario
// La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos
// de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa.
// Necesitaremos realizar las siguientes pruebas en nuestro código:
// Calcular el salario de la categoría “A”.
// Calcular el salario de la categoría “B”.
// Calcular el salario de la categoría “C”.

func TestCategoryA(t *testing.T) {
	//arrange
	var minutesWorkedMonthly float64 = 60
	category := "A"
	var expectedResult float64 = 4500

	//act
	actualResult := calculateSalary(minutesWorkedMonthly, category)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestCategoryB(t *testing.T) {
	//arrange
	var minutesWorkedMonthly float64 = 60
	category := "B"
	var expectedResult float64 = 1800

	//act
	actualResult := calculateSalary(minutesWorkedMonthly, category)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestCategoryC(t *testing.T) {
	//arrange
	var minutesWorkedMonthly float64 = 60
	category := "C"
	var expectedResult float64 = 1000

	//act
	actualResult := calculateSalary(minutesWorkedMonthly, category)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}
