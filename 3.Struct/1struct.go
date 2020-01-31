package main

import "fmt"

type Data struct {  // Collection of  or more fields
	X int
	Y int
}

func main() {
	fmt.Println(Data{1, 2})

	data := Data{1,2}
	data.X = 23
	fmt.Println(data.X)
	
}
