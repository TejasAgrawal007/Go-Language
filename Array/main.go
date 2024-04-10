package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello I am Array!")

	// Create
	var name [5]string

	name[0] = "Tejas"
	name[1] = "Mahesh"
	name[2] = "Sonu"

	// fmt.Println("The name is:", name[0])

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("The Number is: ", numbers)

	// fmt.Println("The length is:", len(numbers))

	/* Go Set the default value

	int  - 0
	bool - False
	String  - "" -> Empty Space
	pointer Complex - nil

	%q -> Queted String


	*/

	// var numbers [5]int
	// fmt.Println("The number is", numbers)

	var data [5]string
	fmt.Println("The String is:", data)
	fmt.Println("The String is %q", data)
}
