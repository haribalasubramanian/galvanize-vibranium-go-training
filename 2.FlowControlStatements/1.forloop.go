package main

import "fmt"

// go has only 1 looping construct- i.e. for loop
// how a typical for loop contains
// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration

// How it is different from other programming languages-- No parenthesis in for loop
// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// init and post are optional
// func main() {
// 	sum := 1
// 	for ; sum < 10; {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

//No while loop in Go. Only for loop
func main(){
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	//infinite loop

	// for{
	// 	fmt.Println("Platform brigade")
	// }
}


//infinite loop
// for {

// }