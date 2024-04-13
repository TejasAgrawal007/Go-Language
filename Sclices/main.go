package main

import (
	"fmt"
)

func main() {
	fmt.Println("Alternative - Dynamic Array")

	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 6, 7, 8, 9, 10)
	// fmt.Println("Number:", numbers)
	// fmt.Println("Data Type = %T\n", numbers)
	// fmt.Println("The length ", len(numbers))

	// name := []string{}
	// fmt.Println("Name:", name)

	// number := []int{1, 2, 3}

	// fmt.Println("Sclice:", number)
	// fmt.Println("Length:", len(number))
	// fmt.Println("Capacity:", cap(number))

	// Make Methods

	number := make([]int, 3, 5)
	// It will take two parameter the first one 3 is a length and the secound one 5 is an capacity of an sclice array...
	// Note : The Length 3 is an already stored size ✅
	// IF The Requrement get increase the capcity is double 🤯

	number = append(number, 4, 5)

	fmt.Println("Sclice:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))
}
