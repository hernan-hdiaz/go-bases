package main

import "fmt"

// Ejercicio 2 - Productos
// Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos
// y retornar el valor del precio total.
// La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

// Y los costos adicionales son:
// Pequeño: solo tiene el costo del producto
// Mediano: el precio del producto + un 3% de mantenerlo en la tienda
// Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500
// de costo de envío.

// El porcentaje de mantenerlo en la tienda es en base al precio del producto.
// El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

// Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz
// Producto que tenga el método Precio.

// Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo
// del producto y los adicionales en caso que los tenga

const (
	smallType  = "small"
	mediumType = "medium"
	largeType  = "large"
)

var additionalCostRate = map[string]float64{
	"small":  1,
	"medium": 1.03,
	"large":  1.06,
}

type product interface {
	price() float64
}

type smallProduct struct {
	BasePrice float64 `json:"base_price"`
}

type mediumProduct struct {
	BasePrice float64 `json:"base_price"`
}

type largeProduct struct {
	BasePrice float64 `json:"base_price"`
}

func main() {
	sp := newProduct(smallType, 100)
	mp := newProduct(mediumType, 100)
	lp := newProduct(largeType, 100)

	fmt.Println("Total price for a small product that costs $100 is:", sp.price())
	fmt.Println("Total price for a medium product that costs $100 is:", mp.price())
	fmt.Println("Total price for a large product that costs $100 is:", lp.price())
}

func newProduct(productType string, basePrice float64) product {
	switch productType {
	case smallType:
		return smallProduct{BasePrice: basePrice}
	case mediumType:
		return mediumProduct{BasePrice: basePrice}
	case largeType:
		return largeProduct{BasePrice: basePrice}
	}
	return nil
}

func (s smallProduct) price() float64 {
	return s.BasePrice * additionalCostRate[smallType]
}

func (m mediumProduct) price() float64 {
	return m.BasePrice * additionalCostRate[mediumType]
}

func (l largeProduct) price() float64 {
	return l.BasePrice*additionalCostRate[largeType] + 2500
}
