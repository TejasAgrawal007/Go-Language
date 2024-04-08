package main

import "fmt"

func firstFunction() {
	fmt.Println("The First Function!")
}

// func addTwoNum(a, b int) int {
// 	return a + b
// }

// Another Approch -
func addTwoNum(a, b int) (result int) {
	result = a + b
	return
}

func mulTwoNum(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are Learning go in go Language!")
	firstFunction()

	result := addTwoNum(10, 40)
	fmt.Println("The Final Output = ", result)

	data := mulTwoNum(10, 20)
	fmt.Println("The Multiply of Two Number is:", data)
}
