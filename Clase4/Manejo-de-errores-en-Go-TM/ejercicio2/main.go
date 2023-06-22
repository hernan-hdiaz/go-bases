package main

import (
	"errors"
	"fmt"
)

// Ejercicio 2 - Impuestos de salario #2
// En tu función “main”, definí una variable llamada “salary” y asignale un valor de
// tipo “int”.
// Creá un error personalizado con un struct que implemente “Error()” con el mensaje
// “Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor
// o igual a  10.000. La validación debe ser hecha con la función “Is()”
// dentro del “main”.

type MyError struct {
}

var err1 = errors.New("el salario es menor a 10.000")

func (e *MyError) Error() error {
	return err1
}

func main() {
	salary := 100
	if salary <= 10000 {
		e := MyError{}

		if errors.Is(e.Error(), err1) {
			fmt.Println("Error:", e.Error())
		}

	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
