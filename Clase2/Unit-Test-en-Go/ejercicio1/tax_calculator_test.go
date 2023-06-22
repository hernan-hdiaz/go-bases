package main

// Ejercicio 1 - Testear el impuesto del salario
// La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados,
// al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos
// de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test
// correspondientes para:
// Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
// Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
// Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTax(t *testing.T) {
	//arrange
	var salary float64 = 200000
	var expectedResult float64 = 54000

	//act
	actualResult := salaryTax(salary)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestBasicTax(t *testing.T) {
	//arrange
	var salary float64 = 100000
	var expectedResult float64 = 17000

	//act
	actualResult := salaryTax(salary)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestNoTax(t *testing.T) {
	//arrange
	var salary float64 = 10000
	var expectedResult float64 = 0

	//act
	actualResult := salaryTax(salary)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}
