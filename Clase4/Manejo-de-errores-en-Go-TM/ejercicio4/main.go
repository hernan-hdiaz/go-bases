package main

import (
	"errors"
	"fmt"
)

// Ejercicio 4 - Impuestos de salario #4
// Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
// para que el mensaje de error reciba por parámetro el valor de “salary”
// indicando que no alcanza el mínimo imponible (el mensaje mostrado por
// consola deberá decir: “Error: el mínimo imponible es de 150.000 y el
// salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int
// pasado por parámetro).

var err1 = errors.New("el mínimo imponible es de 150.000 y el salario ingresado es de")

func x(salary int) error {
	return fmt.Errorf("%w: %d", err1, salary)
}

func main() {
	salary := 2000
	if salary <= 10000 {
		e := x(salary)

		if errors.Is(e, err1) {
			fmt.Println("Error:", e)
		}

	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
