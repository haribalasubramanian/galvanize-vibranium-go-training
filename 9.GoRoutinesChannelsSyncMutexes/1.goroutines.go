package main

import (
	"fmt"
	"time"
)

func printKaro(kya string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(kya)
	}
}

func main() {
	go printKaro("world")
	printKaro("hello")
	
	//printKaro("world")
}	


//A goroutine is a lightweight thread managed by the Go runtime.

