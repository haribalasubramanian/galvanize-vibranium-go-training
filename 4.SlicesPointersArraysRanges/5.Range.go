package main
import "fmt"

var platformBrigade = []string{
	"Suresh","Hari","Soumya", "Asutosh",
	
}

func main(){
	for  i:= range platformBrigade{ // i means Index and v means element at the index
		//fmt.Printf("%v. %v \n", i, v)		
		fmt.Printf("%v\n", i)		
	}
}