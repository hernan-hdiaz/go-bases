package main

// Ejercicio 5 - Calcular cantidad de alimento
// El refugio de animales envió una queja ya que el cálculo total de alimento a comprar
// no fue el correcto y compraron menos alimento del que necesitaban.
// Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios
// para verificar que todo funcione correctamente:
// Verificar el cálculo de la cantidad de alimento para los perros.
// Verificar el cálculo de la cantidad de alimento para los gatos.
// Verificar el cálculo de la cantidad de alimento para los hamster.
// Verificar el cálculo de la cantidad de alimento para las tarántulas.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {
	//arrange
	animalType := dog
	var expectedResult float64 = 100

	//act
	animalDog, _ := animal(animalType)
	actualResult := animalDog(10)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestCat(t *testing.T) {
	//arrange
	animalType := cat
	var expectedResult float64 = 50

	//act
	animalCat, _ := animal(animalType)
	actualResult := animalCat(10)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestHamster(t *testing.T) {
	//arrange
	animalType := hamster
	var expectedResult float64 = 2.5

	//act
	animalHamster, _ := animal(animalType)
	actualResult := animalHamster(10)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}

func TestTarantula(t *testing.T) {
	//arrange
	animalType := tarantula
	var expectedResult float64 = 1.5

	//act
	animalTarantula, _ := animal(animalType)
	actualResult := animalTarantula(10)

	//assert
	assert.Equal(t, expectedResult, actualResult, "must be equals")
}
