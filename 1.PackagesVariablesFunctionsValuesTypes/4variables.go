package main

import (
	"fmt"
	"reflect"
)

var c, python, java bool

var i, j int = 1, 2 //variables initialized. // if a variable is initiliazed , type can be omiited

func main() {
	var i int
	//short variable declarations
	k := 3
	var jk int = 3
	// This should be inside a function
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(reflect.DeepEqual(jk, k))
	fmt.Println(i, c, python, java, k)
}

// Add a function with the function concept
