package main

import "fmt"

func compute(fn func(int,int) int) int{ 	
	return fn(8,4)
}

// func sum1(x,y int) int{
// 	return  x + y
// }

// func computeOutput() int{
// 	return fn(3,4)
// }

//arguments: fn func(int,int) int
//func funcname(arguments) return_type{
	//return ouput with return_type
//}

\
func main(){
	sum := func(x,y int)int{
		return x+ y
	}	

	fmt.Println(sum(2,3))
	fmt.Println(compute(sum))
}