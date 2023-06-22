package main

import (
	"fmt"
	"strconv"
)

// Ejercicio 2 - Calcular promedio

// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y
// devuelva el promedio. No se pueden introducir notas negativas.

func main() {
	fmt.Println("Average is: " + calculateAverage(10, 10, 1))
}

func calculateAverage(grades ...int) (averageIs string) {
	var sum int
	for _, grade := range grades {
		if grade < 0 {
			averageIs = "Grades can't be negative"
			return
		} else {
			sum += grade
		}
	}
	averageNumber := sum / len(grades)
	averageIs = strconv.Itoa(averageNumber)
	return
}
