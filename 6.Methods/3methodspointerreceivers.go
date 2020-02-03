package main

import (
	"fmt" 
)
	
type Values struct {A,B int}


///Sum sums he values
func (v Values) Sum()int {
	fmt.Println(v)
	return (v.B + v.A)
}


//ModifyValues modifis the vaues-- There is a lot of difference with and without a pointer receiver
func (v *Values) ModifyValues(i int) {
	v.A = v.A * i
	v.B = v.B * i
	//fmt.Println(v)
}

func main() {
	v:= Values{3, 4} // This can be also v:= &Values{3,4} Go interprets both as same as the reciever is a pointer.
	v.ModifyValues(10)
	fmt.Println(v.Sum())
}

// /Since methods ofte eed to modify their receiver, pointer receivers are more common than value receivers.
// /ith a value receiver the ModifyValues method operates on a copy of the original Vertex value. (This is th same behavior as for any other function argument.)
// /Choose a value or a pointer receiver



