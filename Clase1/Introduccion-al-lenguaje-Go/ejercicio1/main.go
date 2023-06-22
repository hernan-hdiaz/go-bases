package main

// Ejercicio 1 - Imprimí tu nombre
// Tendrás que:
// Crear una aplicación donde tengas como variable tu nombre y dirección.
// Imprimir en consola el valor de cada una de las variables.

import "fmt"

var (
	name    = "Hernán"
	address = "De la Serna 1635"
)

func main() {
	fmt.Println("Nombre:", name)
	fmt.Println("Dirección:", address)
}
