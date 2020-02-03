package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


//Providing length as 2nd argument to make makes it buffered

//Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

