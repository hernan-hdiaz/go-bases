package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ejercicio 1 - Datos de clientes
// Un estudio contable necesita acceder a los datos de sus empleados para poder realizar
// distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt.
// Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica
// el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa.
// Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt”
// (recordá lo visto sobre el pkg “os”).
// Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el
// programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el
// mensaje “el archivo indicado no fue encontrado o está dañado”.
// Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

// Ejercicio 2 - Imprimir datos

// A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
// Ahora que el archivo sí existe, el panic no debe ser lanzado.
// Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
// Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos
// que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
// Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos
// tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos
// al finalizar su uso.

// Ejercicio 3 - Registrando clientes
// El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar
// datos de nuevos clientes. Los datos requeridos son:
// Legajo
// Nombre
// DNI
// Número de teléfono
// Domicilio

// Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello,
// necesitás leer los datos de un array. En caso de que esté repetido, debes manipular
// adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
// 1.- generar un panic;
// 2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución
// del programa normalmente.

// Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función
// para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero.
// Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error
// para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada
// 	tipo de dato, ej: 0, “”, nil).

// Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por
// consola los siguientes mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo
// de ejecución”. Utilizá defer para cumplir con este requerimiento.

// Requerimientos generales:
// Utilizá recover para recuperar el valor de los panics que puedan surgir
// Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
// Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go
// (realiza también la validación pertinente para el caso de error retornado).

// Legajo
// Nombre
// DNI
// Número de teléfono
// Domicilio

var (
	ErrCanNotCreate          = errors.New("can not create file")
	ErrCanNotOpen            = errors.New("can not open file")
	ErrCanNotRead            = errors.New("can not read file")
	ErrCanNotWrite           = errors.New("can not write file")
	ErrCustomerAlreadyExists = errors.New("customer DNI already exists")
	newFileName              = "customers.txt"
	errorsOcurred            = false
)

type Customer struct {
	Legajo    string `json:"legajo"`
	Nombre    string `json:"nombre"`
	DNI       string `json:"dni"`
	Telefono  string `json:"telefono"`
	Domicilio string `json:"domicilio"`
}

func main() {
	customer :=
		Customer{
			Legajo: "147",
			Nombre: "Pedro Picapiedras",
			// DNI:    "121314150",
			DNI:       "12131415",
			Telefono:  "40000001",
			Domicilio: "Camino de piedra 1",
			// Domicilio: "",
		}
	createFile(newFileName)
	printFile(newFileName)
	addCustomer(newFileName, customer)
	// addCustomer("nonExistingFileName.txt", customer)
	printFile(newFileName)
	if errorsOcurred {
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}
	fmt.Println("Fin de la ejecución")
}

func addCustomer(fileName string, customer Customer) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Error:", err)
			errorsOcurred = true
		}
	}()
	if !customer.Exists(newFileName) {

		validFields, err := validateFields(customer)
		if err != nil {
			panic(err)
		}

		if validFields {
			file, err := os.OpenFile("./"+fileName, os.O_APPEND|os.O_RDWR, os.ModeAppend)
			if err != nil {
				panic("el archivo indicado no fue encontrado o está dañado")
			}

			_, err3 := file.WriteString("\n" +
				customer.Legajo + "," +
				customer.Nombre + "," +
				customer.DNI + "," +
				customer.Telefono + "," +
				customer.Domicilio)
			if err3 != nil {
				panic(err3)
			}
			file.Close()
		}
	}
}

func validateFields(customer Customer) (bool, error) {
	switch "" {
	case customer.Legajo:
		return false, errors.New("se ingresó por parámetro un Legajo vacío")
	case customer.Nombre:
		return false, errors.New("se ingresó por parámetro un Nombre vacío")
	case customer.DNI:
		return false, errors.New("se ingresó por parámetro un DNI vacío")
	case customer.Telefono:
		return false, errors.New("se ingresó por parámetro un Número de Teléfono vacío")
	case customer.Domicilio:
		return false, errors.New("se ingresó por parámetro un Domicilio vacío")
	}
	return true, nil
}

func (c *Customer) Exists(fileName string) (exist bool) {
	file := readFile(fileName)

	str1 := string(file[:])

	r := csv.NewReader(strings.NewReader(str1))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Error:", err)
			errorsOcurred = true
		}

	}()
	for _, row := range records {
		for i, v := range row {
			if i == 2 {
				if v == c.DNI {
					exist = true
					panic("el cliente ya existe")
				}
			}
		}
	}
	return
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(ErrCanNotCreate)
	}
	defer file.Close()
	_, err2 := file.WriteString("123,Juan Perez,12345678,12345678,Falso 123\n" +
		"456,Pedro Robles,23242526,12345678,Falso 123\n" +
		"789,Camilo Sexto,12131415,12345678,Falso 123")
	if err2 != nil {
		log.Fatal(ErrCanNotWrite)
	}
}

func printFile(fileName string) {
	file := readFile(fileName)
	fmt.Printf("%s\n", file)
}

func readFile(fileName string) []byte {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Error:", err)
			errorsOcurred = true
		}

	}()
	file, err := os.ReadFile("./" + fileName)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return file
}
