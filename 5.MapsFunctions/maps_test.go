package main

import (
	"testing"
	"fmt"
)

//Format of a testing function
// func TestXXXXXXXXXXXXXXXXXX(t *testing.T){

// }
func TestIfKeyExists(t *testing.T){
	var learningMap map[string]Employee
	learningMap = make(map[string]Employee)

	learningMap["Soumya"] = Employee{
		27, 560068,"JP NAGAR",
	}

	learningMap["Hari"] = Employee{
		24, 560088,"Whitefiled",
	}

	exists := CheckIfKeyExists("Hari",learningMap)
	if exists{
		//Success
		fmt.Printf("Test Case 0: Success, Expected output and output recieved is same.")
	}else{
		//Error
		t.Errorf("Test Case 0: Expected output is %v but got %v",true ,exists)
	}
	delete(learningMap, "Hari")
	exists = CheckIfKeyExists("Hari",learningMap)
	if !exists{
		//Success
		fmt.Printf("Test Case 1: Success, Expected output and output recieved is same.")
	}else{
		//Error
		t.Errorf("Test Case 1: Expected output is %v but got %v",true ,exists)
	}

	//Algo
	

}


// func TestIfKeyExistsOrNot(t *testing.T){
// 	//1. We have to create a dummy Map. and a key..
// 	var testMap map[string]Employee
// 	testMap = make(map[string]Employee)

// 	testMap["Asu"] = Employee{
// 		27,
// 		560078,
// 		"Marathahali",
// 	}

// 	testMap["Biswa"] = Employee{
// 		28,
// 		560068,
// 		"JPNagar",
// 	}

// 	//Biswa and Asu
// 	isKeyExists := CheckIfKeyExists("Biswa1", testMap)
// 	if !isKeyExists{
// 		t.Errorf("Test Case 0: Expected output is %v but got %v",true ,isKeyExists)
// 	}else{
// 		fmt.Printf("Test Case 0: Success, Expected output and output recieved is same.")
// 	}

// 	delete(testMap, "Biswa")

// 	isKeyExists = CheckIfKeyExists("Biswa", testMap)
// 	if !isKeyExists{
// 		t.Errorf("Test Case 1: Expected output is %v but got %v",true ,isKeyExists)
// 	}else{
// 		fmt.Printf("Test Case 1: Success, Expected output and output recieved is same.")
// 	}


// }

// //Coverage commands needs to keep in mind
// go test -cover
// go test -coverprofile=coverage.out 	//The -coverprofile flag automatically sets -cover to enable coverage analysis
// go tool cover -func=coverage.out  //we can ask for the coverage to be broken down by function, although that's not going to illuminate much in this case since there's only one function
// go tool cover -html=coverage.out // html 

