package main

import "fmt"

// Ejercicio 2 - Préstamo

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
// Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
// Solo le otorga préstamos a clientes
// cuya edad sea mayor a 22 años
// se encuentren empleados y
// tengan más de un año de antigüedad en su trabajo.
// Dentro de los préstamos que otorga no les cobrará
// interés a los que posean un sueldo superior a $100.000.
// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de
// acuerdo a cada caso.
// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.

const (
	approved_age         int     = 22
	approved_laborOld    int     = 1
	interest_free_salary float32 = 100000
)

var (
	age      int     = 22
	employed bool    = true
	laborOld int     = 1
	salary   float32 = 100001
)

func main() {
	if !(age >= approved_age) {
		fmt.Println("Préstamo rechazado, debe ser mayor de", approved_age, "años")
	} else if !employed {
		fmt.Println("Préstamo rechazado, debe encontrarse empleado")
	} else if !(laborOld >= approved_laborOld) {
		fmt.Println("Préstamo rechazado, debe tener más de", approved_laborOld, "año de antigüedad en su trabajo")
	} else {
		switch {
		case salary > interest_free_salary:
			fmt.Println("Préstamo otorgado sin intereses")
		case salary > 0:
			fmt.Println("Préstamo otorgado con intereses")
		default:
			fmt.Println("El salario debe ser mayor a 0, revise la información y vuelva a intentarlo")
		}
	}
}
