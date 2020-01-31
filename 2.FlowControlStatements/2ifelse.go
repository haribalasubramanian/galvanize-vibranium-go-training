package main

import (
	"fmt"
	
)

func isPositive(x int) string {
	// no paranthesis . only brackets required. same as for loop
	if x < 0 {
		return "negative"
	}
	return "positive"
}

func isDivisibleBy2(x int) string{
	// if with short statement
	if h := x %2; h == 0{
		return "Divisible"
	}
	return "not divisible"
}

func ifelseExample(name string) string{
	if name == "banu"{
		return "Platform brigade"
	}else{
		return "No idea."
	}
}
func main() {
	fmt.Println(isPositive(2))
	fmt.Println(isPositive(-1))

	fmt.Println(isDivisibleBy2(2))

	fmt.Println("Find brigade name")
	fmt.Println(ifelseExample("banu"))
	fmt.Println(ifelseExample("xys"))
}
