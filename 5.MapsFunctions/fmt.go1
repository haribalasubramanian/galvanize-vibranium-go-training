package main


import "fmt"

func main(){
	var name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)	//func Errorf(format string, a ...interface{}) error
	fmt.Println(err.Error())

	var z="asdasd"
	var y = "asd"
	"User entered " + z + "but I got output as " + y
	fmt.Printf("User entered %v but I got output as %v" , z, y)
	fmt.XXXXf("User entered z but I got output as y")
	%v - ddefault format
	%p - pointer and structs
	%T- Type 
	%s - 'string'
	%f - Floating Point
	%d - integer
	%t - boolean

	fmt.Printf("User entered %v but I got output as %v" , z, y)
	fmt.Println("We are platform brigade")

	name = "Kim"
	age :=  22
	fmt.Printf("%s is %d years old.\n", name, age)	//func Printf(format string, a ...interface{}) (n int, err error) .. Not to worry about error returned by printf
	fmt.Println(name, "is", age, "years old.")		//	func Println(a ...interface{}) (n int, err error)

	w :=fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Println(w)

	var s1   string
	var s2   string
	var rint int
	
	_, err = fmt.Scanf("%s %d %s", &s1, &rint, &s2)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("this is the first string from the commandline: ", s1)
	fmt.Println("this is the first int from the commandline: ", rint)
	fmt.Println("this is the second string from the commandline: ", s2)
	
}