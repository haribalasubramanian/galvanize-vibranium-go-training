package main

import "fmt"

func main() {
	i:= 42

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i


}

// Go does not supports pointer arithmetic.

// slices: A pointer need to be used when you may need to modify the length or capacity, which changes the value of the slice.
// maps: Don't use a pointer, since the map value doesn't change with modifications.
// functions and methods: Don't use a pointer, the same effect is had through function values.
// chan: Don't use a pointer.


// Go has pointers. A pointer holds the memory address of a value.

// The type *T is a pointer to a T value. Its zero value is nil.

