package main

import "fmt"

// // Ejercicio 3 - A qué mes corresponde

// // Realizar una aplicación que reciba  una variable con el número del mes.
// // Según el número, imprimir el mes que corresponda en texto.
// // ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
// // Ej: 7, Julio.
// // Nota: Validar que el número del mes, sea correcto.

var (
	months = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	monthNumber int = 12
)

func main() {
	if monthNumber < 1 || monthNumber > 12 {
		fmt.Println("El número de mes debe ser entre 1 y 12 inclusive")
	} else {
		fmt.Printf("%d, %s\n", monthNumber, months[monthNumber])
	}
}
