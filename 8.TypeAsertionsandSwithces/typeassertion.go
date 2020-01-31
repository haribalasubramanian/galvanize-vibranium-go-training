package main

import "fmt"

func main() {
	var iface interface{} = "Vibranium Brigade"

	value := iface.(string)
	fmt.Println(value)

	value, ok := iface.(string)
	fmt.Println(value, ok)

	floatValue, ok := iface.(float64)
	fmt.Println(floatValue, ok)

	floatValue = iface.(float64) // panic
	fmt.Println(floatValue)
}

// A type assertion provides access to an interface value's underlying concrete value.

// t := i.(T)