package main

import "fmt"

// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad
// para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables
// Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	a := Alumnos{
		Nombre:   "Juan",
		Apellido: "Pérez",
		DNI:      39456789,
		Fecha:    "20-03-2023",
	}

	a.Detalle()
}

func (a Alumnos) Detalle() {
	fmt.Printf(
		"Nombre: %s\n"+
			"Apellido: %s\n"+
			"DNI: %d\n"+
			"Fecha: %s\n",
		a.Nombre,
		a.Apellido,
		a.DNI,
		a.Fecha)
}
