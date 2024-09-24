package main

import (
	"fatory/products"
	"fmt"
)

func main() {
	factory := products.Product{}
	product := factory.New("My Product")

	fmt.Printf("My product is %+v\n", product)
}
