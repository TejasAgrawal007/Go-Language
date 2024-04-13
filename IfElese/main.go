package main

import "fmt"

func main() {
	// var x int
	// fmt.Println("Enter Your Numebr: ")

	// fmt.Scan(&x)

	// if x > 0 {
	// 	fmt.Println("Positive Number")
	// } else if x == 0 {
	// 	fmt.Println("Number is Netural")
	// } else {
	// 	fmt.Println("Negative Number")
	// }

	var age int
	fmt.Println("Enter Your Age: ")
	fmt.Scan(&age)

	if age >= 18 && age <= 90 {
		fmt.Println("You Can Work With Us")
	} else {
		fmt.Println("You Can't work with Us")
	}
}
