package main

import (
	"fmt"
	
)

//Values structs the values needed to calculate the sum
type Values struct {
	A, B int
}

//Pythagoreantheorem calculates the pythagorean theorem 
func (v Values) Pythagoreantheorem() int { //Pythagoreantheorem method has a reciever of type Values named v
	return (v.A * v.A + v.B * v.B)
}

func main() {
	v := Values{3, 4}
	fmt.Println(v.Pythagoreantheorem())
}


// Techincal Tips:
//Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.
//

//Important: Method is a function but with a receiver argument,

