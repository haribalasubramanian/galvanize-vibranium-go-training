package main

import "fmt"

func main() {
	var we interface{}
	getDetails(we)

	we = 21
	getDetails(we)

	we = "We are a brand."
	getDetails(we)
}

func getDetails(we interface{}) {
	fmt.Printf("(%v, %T)\n", we, we)
}
