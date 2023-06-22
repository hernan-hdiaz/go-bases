package main

import "fmt"

// Ejercicio 2 - Employee
// Una empresa necesita realizar una buena gestión de sus empleados,
// para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente
// dichos empleados. Los objetivos son:
// Crear una estructura Person con los campos ID, Name, DateOfBirth.
// Crear una estructura Employee con los campos: ID, Position y una composicion con la
// estructura Person.
// Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará
// es realizar la impresión de los campos de un empleado.
// Instanciar en la función main() tanto una Person como un Employee cargando sus
// respectivos campos y por último ejecutar el método PrintEmployee().
// Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar
// la gestión de los empleados.

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

type Employee struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Person   `json:"person"`
}

func main() {
	p := Person{
		ID:          2,
		Name:        "Juan",
		DateOfBirth: "20-05-1990",
	}

	e := Employee{
		ID:       1,
		Position: "Boss",
		Person:   p,
	}

	e.PrintEmployee()
}

func (e Employee) PrintEmployee() {
	fmt.Printf(
		"ID: %d\n"+
			"Position: %s\n"+
			"PersonID: %d\n"+
			"Name: %s\n"+
			"DateOfBirth: : %s\n",
		e.ID,
		e.Position,
		e.Person.ID,
		e.Name,
		e.DateOfBirth)
}
