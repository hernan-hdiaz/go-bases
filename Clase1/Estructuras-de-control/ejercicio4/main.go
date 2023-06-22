package main

import "fmt"

// Ejercicio 4 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
// Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

func main() {
	employeeName := "Benjamin"
	fmt.Printf("Empleado: %s, Edad: %d\n", employeeName, employees[employeeName])

	var employeesOver21 int
	for _, age := range employees {
		if age >= 21 {
			employeesOver21++
		}
	}
	fmt.Printf("La cantidad de empleados mayores a 21 años es %d\n", employeesOver21)

	fmt.Println(employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
