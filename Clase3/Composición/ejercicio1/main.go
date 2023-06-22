package main

import "fmt"

// Ejercicio 1
// Crear un programa que cumpla los siguiente puntos:
// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products instanciado con valores.
// 2 métodos asociados a la estructura Product: Save(), GetAll().
// El método Save() deberá tomar el slice de Products y añadir el producto desde el cual
// se llama al método.
// El método GetAll() deberá imprimir todos los productos guardados
// en el slice Products.
// Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el
// producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definido desde main().

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Pantene",
		Price:       100,
		Description: "Shampoo para el cabello",
		Category:    "Cuidado personal",
	},
	{
		ID:          2,
		Name:        "Dove",
		Price:       150,
		Description: "Barra de belleza corporal",
		Category:    "Cuidado personal",
	},
	{
		ID:          3,
		Name:        "Galletitas Paseo 5 semillas",
		Price:       85,
		Description: "Galletas de masa madre con 5 semillas",
		Category:    "Alimentos",
	},
}

func main() {
	p := Product{
		ID:          4,
		Name:        "Galletitas Oreo",
		Price:       120,
		Description: "Galletas de chocolate rellenas de pasta sabor vainilla",
		Category:    "Alimentos",
	}
	p.Save()
	p.GetAll()
	fmt.Println(getById(2))
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf(
			"ID: %d\n"+
				"Name: %s\n"+
				"Price: %.2f\n"+
				"Description: %s\n"+
				"Category: : %s\n",
			product.ID,
			product.Name,
			product.Price,
			product.Description,
			product.Category)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if id == product.ID {
			return product
		}
	}
	return Product{}
}
