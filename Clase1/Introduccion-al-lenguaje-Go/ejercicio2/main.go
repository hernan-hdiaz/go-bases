package main

// Ejercicio 2 - Clima

// Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
// Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprimí los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

import "fmt"

var (
	temperature = 12
	humidity    = 56
	pressure    = 1030
)

func main() {
	fmt.Println("Temperatura:", temperature)
	fmt.Println("Humedad:", humidity)
	fmt.Println("Presión:", pressure)
}
