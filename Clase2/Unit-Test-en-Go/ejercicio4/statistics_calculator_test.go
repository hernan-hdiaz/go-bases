package main

// Ejercicio 4 - Testear el cálculo de estadísticas
// Los profesores de la universidad de Colombia, entraron al programa de análisis de datos
// de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores
// nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos.
// Se solicita la siguiente tarea:
// Realizar test para calcular el mínimo de calificaciones.
// Realizar test para calcular el máximo de calificaciones.
// Realizar test para calcular el promedio de calificaciones.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	//arrange
	operationType := minimum
	expectedResult := 1

	//act
	minFunc, _ := operation(operationType)
	actualResult := minFunc(10, 5, 1, 3, 5)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestMax(t *testing.T) {
	//arrange
	operationType := maximum
	expectedResult := 10

	//act
	maxFunc, _ := operation(operationType)
	actualResult := maxFunc(10, 5, 1, 3, 5)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestAverage(t *testing.T) {
	//arrange
	operationType := average
	expectedResult := 4

	//act
	averageFunc, _ := operation(operationType)
	actualResult := averageFunc(10, 5, 1, 3, 5)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}
