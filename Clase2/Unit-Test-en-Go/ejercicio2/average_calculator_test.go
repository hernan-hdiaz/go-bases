package main

// Ejercicio 2 - Calcular promedio
// El colegio informó que las operaciones para calcular el promedio no se están realizando
// correctamente, por lo que ahora nos corresponde realizar los test correspondientes:
// Calcular el promedio de las notas de los alumnos.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegativeGrades(t *testing.T) {
	//arrange
	expectedResult := "Grades can't be negative"

	//act
	actualResult := calculateAverage(10, 5, -5, 20)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestAverage(t *testing.T) {
	//arrange
	expectedResult := "7"

	//act
	actualResult := calculateAverage(10, 10, 1)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}
