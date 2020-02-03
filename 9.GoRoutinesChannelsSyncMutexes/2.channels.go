package main

import "fmt"

func sum(s []int, channel chan int) {
	sum := 0
	for _, v := range s {
		sum += v

	}
	
	//c <- sum // send sum to c

	channel <- sum
}

func main() {
	me := []int{10,9,8,7,6}
	//you := []int{5,4,3,2,1}
	channel := make(chan int) // Syntax to create channels.
	
	go sum(me,channel)
	i := <-channel
	fmt.Println(i)
	//go sum(you,channel)
	//i,u := <-channel,<-channel //receives a channel
//	fmt.Println(i + u)	
}



// Send a data to a channel

// channel <- data

// data := <- channel
