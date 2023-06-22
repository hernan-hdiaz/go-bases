package main

import "fmt"

// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener
// cada una de las letras por separado para deletrearla. Para eso tendrán que:
// Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras
// que contiene la misma.
// Luego, imprimí cada una de las letras.

func main() {
	word := "Cinco"
	letterQuantity := len(word)
	fmt.Println(letterQuantity)

	for i := 0; i < letterQuantity; i++ {
		fmt.Println(string(word[i]))
	}
}
