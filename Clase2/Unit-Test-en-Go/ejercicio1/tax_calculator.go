package main

import "fmt"

// Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
// para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y
// si gana más de $150.000 se le descontará además un 10 % (27% en total).

var salary float64 = 200000

const (
	basicTax         = 17.00
	complementaryTax = 10.00
)

func main() {
	fmt.Println("El impuesto para el salario de $", salary, "es de $", salaryTax(salary))
}

func salaryTax(salary float64) (salaryAfterTaxes float64) {
	switch {
	case salary > 150000:
		salaryAfterTaxes = salary * (basicTax + complementaryTax) / 100
	case salary > 50000:
		salaryAfterTaxes = salary * basicTax / 100
	}
	return
}
