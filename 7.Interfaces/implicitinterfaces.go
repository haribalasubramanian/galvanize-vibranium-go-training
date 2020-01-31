package main

import "fmt"

//Galvanites Used as an interface
type Galvanites interface {
	PrintNames()
}


//Vibranium structs the 
type Vibranium struct {
	MemberName string
}

// This method means type Vibranium implements the interface Galvanites,
// No need to explicitly declare that it does
func (v Vibranium) PrintNames() {
	fmt.Println(v.MemberName)
}

func main() {
	var we Galvanites = Vibranium{"Suresh"}
	we.PrintNames()
}
