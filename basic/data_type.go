package main

import (
	"fmt"
	"reflect"
)

func main() {

	// data type
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(42.2))
	fmt.Println(reflect.TypeOf("42"))
	fmt.Println(reflect.TypeOf(true))
	// variable declaration : fixed or static data type
	// One way - > var variable_name type

	var quantity int = 2
	var length, width float64 = 1.2, 3.4
	var customerName string = "Avijeet Shil"

	// 2nd way +> variable_name :=

	quantity := 2
	length, width := 1.2, 3.4

	customerName := "Avijeet Shil"
	fmt.Println(reflect.TypeOf(quantity))

	fmt.Println(reflect.TypeOf(length))

	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(customerName))

	// type-casting /conversion
	var price int = 100
	var taxRate float64 = .09
	var tax float64 = float64(price) * taxRate // conversion of type from int to float

}
