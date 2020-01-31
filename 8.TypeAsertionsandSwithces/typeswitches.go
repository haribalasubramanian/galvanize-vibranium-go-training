package main

import "fmt"

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Value is %v\n", v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("Value is  %v ", v)
	default:
		fmt.Printf("No idea of type of value entered", v)
	}
}

func main() {
	getType(21)
	getType("hello")
	getType(true)
}


//A type switch is a construct that permits several type assertions in series.

