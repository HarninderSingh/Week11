package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func main() {
	fmt.Println("Addition is ", Add(2, 5))
	fmt.Println("Subtracttion is ", Subtract(2, 5))
	fmt.Println("Multiplication is ", Multiply(2, 5))

}
