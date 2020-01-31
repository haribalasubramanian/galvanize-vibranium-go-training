package main

import "fmt"
// array is of fixed size. but slice is of dynamic sized.

func main(){
	// Slice is about two indices : low bound and high bound. emp[low:high]
	//START: Slices
			// numbers := [6]int{2, 3, 5, 7, 11, 13}
			// var subsets []int = numbers[1:4]
			// fmt.Println(subsets)
	//END: Slices

	// Slices are like reference to arrays.
	// Any change made to a slice modifies the corresponding elements of its underlying array.
	// Other slices that use the underlying array will also see the change.
	// START: Reference to Arrays
	// 		platformbrigade := [6]string{
	// 			"Suresh",
	// 			"Uthra",
	// 			"Soumya",
	// 			"Pooja1",
	// 			"Pooja2",
	// 			"Hari",
	// 		}
	// 		fmt.Println(platformbrigade)

	// 		x := platformbrigade[0:2]
	// 		y := platformbrigade[1:3]
	// 		fmt.Println(x, y)

	// 		y[0] = "Asutosh"
	// 		fmt.Println(x, y)
	// 		fmt.Println(platformbrigade)
	// // END: Reference to Arrays

	// //The default is zero for the low bound and the length of the slice for the high bound.


	// //START: Slice Defaults
	// nums := []int{2, 3, 5, 7, 11, 13}

	// nums = nums[1:4]

	
	
	// fmt.Println(nums) // 3,5,7
	// fmt.Println("Length and Capacity")
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums = nums[:]
	// fmt.Println(nums) //3,5
	// fmt.Println("Length and Capacity")
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums = nums[1:]
	// fmt.Println(nums)// 5
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))
	// //END: Slice Defaults
		
	// //7 Jan 2020
	// // Slice: Length and Capacity
	// fmt.Println("Length and Capacity")
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//Nil Slice
	var sliceArry []int
	fmt.Println(len(sliceArry))
	fmt.Println(cap(sliceArry))
	fmt.Println(sliceArry)	
	if sliceArry == nil{ 
		fmt.Printf("%s", sliceArry)
	}
	
	//	A nil slice has a length and capacity of 0 and has no underlying array.

}

//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

//The length of a slice is the number of elements it contains.

// This is an example of hwhat has been done for testing the pull request functionality.



