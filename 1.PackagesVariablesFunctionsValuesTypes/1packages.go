package main // program starts running with main package

import (
	"fmt"       // this packages are used
	"math/rand" // package name is same as the last element of the import path
)

func main() {
	// THis generates a random number
	fmt.Println("My favorite number is", rand.Intn(10))
}
