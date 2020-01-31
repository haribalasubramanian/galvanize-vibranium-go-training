package main

import "fmt"

func add(x int, y int) int { // type comes after the variable.
	return x + y
}
func addConsecutive(x, y int) int { // type can be omitted when two or more consecutive params share the same type.
	return x + y
}

func multipleResults(x int, y int) (int, int) { //function can return multiple results also.
	return y, x
}

//Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // A return statement without arguments returns the named return values. This is known as a "naked" return.
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addConsecutive(42, 13))
	fmt.Println(multipleResults(4, 5))
	fmt.Println(split(17))
}
