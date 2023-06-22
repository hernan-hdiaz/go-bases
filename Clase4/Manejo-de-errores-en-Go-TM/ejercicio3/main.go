package main

import (
	"errors"
	"fmt"
)

// Ejercicio 3 - Impuestos de salario #3

// Hacé lo mismo que en el ejercicio anterior pero reformulando
// el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.

type MyError struct {
}

var err1 = errors.New("el salario es menor a 10.000")

func x() error {
	return err1
}

func main() {
	salary := 2000
	if salary <= 10000 {
		e := x()

		if errors.Is(e, err1) {
			fmt.Println("Error:", e)
		}

	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
