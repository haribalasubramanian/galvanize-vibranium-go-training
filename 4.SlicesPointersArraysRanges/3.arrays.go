package main

import "fmt"

func main() {
	// Syntax : The type [n]T is an array of n values of type T.

	var a [2]string // array pf two strings
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13, 34}
	fmt.Println(primes)
	//arrays cannot be resized. -- but go comes up with slices which will help you to achieve the same.
}
