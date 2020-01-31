package main

import (
	"fmt"
	"strings"
)
func main(){
		// You can make a slice too by using make.
		// makeArray := make([]string,5)
		// printLenandCapacity(makeArray)

	


		// makeArrayWithCap := make([]string,0,10)
		// printLenandCapacity(makeArrayWithCap)

		// makeArrayWithCap = append(makeArrayWithCap, "Suresh")
		// printLenandCapacity(makeArrayWithCap)

		// arrryorslice = append(arrayorslice, element, eleent2, elemtn2)
		
		//Slices of Slices
		 pyramid()

		// //Slice append.
		//appendSlice()
}


func printLenandCapacity(s []string){
	fmt.Println("Length")
	fmt.Println(len(s))
	fmt.Println("Capacity")
	fmt.Println(cap(s))
	fmt.Println("Slice")
	fmt.Println(s)
}

func pyramid(){
		// Create a pyramid
		pyramid := [][]string{
			[]string{" ", " ", "*"," "," "},
			[]string{" ", "*", "*", "*", " "},
			[]string{"*", "*", "*","*", "*"},
		}	

		// Slices of Slices


	
			
		
		for i := 0; i < len(pyramid); i++ {
			fmt.Printf("%s\n", strings.Join(pyramid[i], " "))
		}
}


// type Employee struct{
// 	Name string
// }

// type Employees []Employee


func appendSlice(){
	var appendSlice []string
	printLenandCapacity(appendSlice)
	
	appendSlice = append(appendSlice, "Platform Brigade")
	printLenandCapacity(appendSlice)

	
	appendSlice = append(appendSlice, "Suresh")
	printLenandCapacity(appendSlice)

	
	appendSlice = append(appendSlice, "PoojaMA", "PoojaS", "Soumya")
	printLenandCapacity(appendSlice)


	// employees := Employees{}

	// employee1 := Employee{
	// 	"Suresh"
	// }
	// employees21 := Employees{
	// 	{"PoojaMA"},
	// 	{"PoojaS"},
	// 	{"Soumya"},
	// }

	// employees = append(employees, employee)

	// //employees = append(employees, ...employees21)

}