package main

import (
	"errors"
	"fmt"
)

// Ejercicio 4 - Calcular estadísticas

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
// de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus
// calificaciones.
// Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar
// (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido)
// que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
// Ejemplo:

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err.Error())
	}
	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println(err.Error())
	}
	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err.Error())
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)
}

func operation(operationType string) (func(grades ...int) int, error) {
	switch operationType {
	case minimum:
		return calculateMinimum, nil
	case average:
		return calculateAverage, nil
	case maximum:
		return calculateMaximum, nil
	default:
		return nil, errors.New("invalid operation type")
	}
}

func calculateMaximum(grades ...int) int {
	maxNum := grades[0]
	for _, grade := range grades {
		if grade > maxNum {
			maxNum = grade
		}
	}
	return maxNum
}

func calculateMinimum(grades ...int) int {
	minNum := grades[0]
	for _, grade := range grades {
		if grade < minNum {
			minNum = grade
		}
	}
	return minNum
}

func calculateAverage(grades ...int) int {
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	average := sum / len(grades)
	return average
}
