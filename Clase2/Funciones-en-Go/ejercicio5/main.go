package main

import "fmt"

// Ejercicio 5 - Calcular cantidad de alimento

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
// Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan
// darle refugio a muchos animales más.
// Tienen los siguientes datos:
// Perro 10 kg de alimento.
// Gato 5 kg de alimento.
// Hamster 250 g de alimento.
// Tarántula 150 g de alimento.
// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el
// animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
// tipo de animal especificado.
// Ejemplo:

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

var animalFood = map[string]float64{
	"dog":       10,
	"cat":       5,
	"hamster":   0.25,
	"tarantula": 0.15,
}

func main() {
	var amount float64

	animalDog, msg := animal(dog)
	if msg != "" {
		fmt.Println(msg)
	} else {
		amount += animalDog(10)
		fmt.Printf("Food needed: %.2f kg\n", amount)
	}

	animalCat, msg := animal(cat)
	if msg != "" {
		fmt.Println(msg)
	} else {
		amount += animalCat(10)
		fmt.Printf("Food needed: %.2f kg\n", amount)
	}
	animalHamster, msg := animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	} else {
		amount += animalHamster(10)
		fmt.Printf("Food needed: %.2f kg\n", amount)
	}
	animalTarantula, msg := animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	} else {
		amount += animalTarantula(10)
		fmt.Printf("Food needed: %.2f kg\n", amount)
	}

}

func animal(animalType string) (func(dogQuantity float64) (foodQuantity float64), string) {
	switch animalType {
	case dog:
		return dogFood, ""
	case cat:
		return catFood, ""
	case hamster:
		return hamsterFood, ""
	case tarantula:
		return tarantulaFood, ""
	default:
		return nil, "calculation not available for " + animalType
	}
}

func dogFood(dogQuantity float64) (foodQuantity float64) {
	foodQuantity = dogQuantity * animalFood["dog"]
	return
}

func catFood(catQuantity float64) (foodQuantity float64) {
	foodQuantity = catQuantity * animalFood["cat"]
	return
}

func hamsterFood(hamsterQuantity float64) (foodQuantity float64) {
	foodQuantity = hamsterQuantity * animalFood["hamster"]
	return
}

func tarantulaFood(tarantulaQuantity float64) (foodQuantity float64) {
	foodQuantity = tarantulaQuantity * animalFood["tarantula"]
	return
}
