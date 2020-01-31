package main
import (
    "fmt"
)
//Sum2Num is used to calculate sum of 2 numbers
func Sum2Num(n1 int, n2 int) int {
    return n1 + n2
}

//Sum3Num is used to calculate sum of 3 numbers
func Sum3Num(num1, num2, num3 int) int {
    return num1 + num2 + num3
}
func main() {
    fmt.Println(Sum2Num(1, 2))
    fmt.Println(Sum3Num(1, 2, 3))
}