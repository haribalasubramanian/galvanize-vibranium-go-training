package main

import (
	"fmt"
	"math"
)

// This is also valid but multiple factored import looks stylish and is a good style.

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi) // This will show an error. Exported names always start with capital letter.
	//When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
}

