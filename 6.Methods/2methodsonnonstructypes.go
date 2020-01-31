package main

import (
	"fmt"
)

//Me is a int struct
type Me int

//CheckForPostiveInteger checks if the number entered is positive or not
func (m Me) CheckForPostiveInteger() string {
	if m < 0 {
		return "Negative"
	}
	return "Positive"
}

func main() {
	me := Me(6)
	fmt.Println(me.CheckForPostiveInteger())
}
