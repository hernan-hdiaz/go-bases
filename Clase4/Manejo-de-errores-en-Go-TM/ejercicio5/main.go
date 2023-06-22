package main

import (
	"errors"
	"fmt"
)

// Ejercicio 5 -  Impuestos de salario #5
// Vamos a hacer que nuestro programa sea un poco más complejo y útil.
// Desarrollá las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora
// como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le
// deberá descontar el 10 % en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o
// un número negativo, la función debe devolver un error. El mismo tendrá que
// indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

func main() {
	var hoursWorkedMonthly float64 = 10
	var hourValue float64 = 10
	salary, err := calculateSalary(hoursWorkedMonthly, hourValue)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El salario para %.0fhs mensuales trabajadas a $%.2f por hora, es de $%.2f\n", hoursWorkedMonthly, hourValue, salary)
	}
}

func calculateSalary(hoursWorkedMonthly, hourValue float64) (salary float64, err error) {
	if hoursWorkedMonthly < 80 {
		err = errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}
	salary = hoursWorkedMonthly * hourValue
	if salary >= 150000 {
		salary *= 0.9
	}
	return
}
