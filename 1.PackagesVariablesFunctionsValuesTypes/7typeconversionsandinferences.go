package main

import (
	"fmt"
	 "math"
)

func main() {

 	var x, y int = 3, 4
//var f float64 = math.Sqrt(float64(x*x + y*y)) // type conversion

// 	var f float64 = 3.9
 	var z int = int(math.Sqrt(float64(x*x + y*y)))   // explicit type conversion 
	fmt.Println( z)


	// v := 123 
	// fmt.Printf("v is of type %T \n", v)  // %T helps you in inferring the type of a variable.
	
}
