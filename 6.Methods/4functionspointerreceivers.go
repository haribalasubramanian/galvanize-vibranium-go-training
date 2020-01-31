package main

import (
	"fmt"
	
)

//Values structs the values
type Values struct {
	A, B int
}

//Sum sums the values
func (v Values) Sum() int {
	return (v.B + v.A)
}

//ModifyValues modifies the values
func ModifyValues(v *Values,i int) {
	v.A = v.A * i
	v.B = v.B * i
	fmt.Println(v)	
}

func main() {
	v := Values{3, 4}
	ModifyValues(&v,10)
	fmt.Println(v.Sum())
}
