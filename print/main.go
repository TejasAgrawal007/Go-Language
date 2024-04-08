package main

import "fmt"

func main() {
	age := 22
	name := "Tejas"
	height := 5.82

	fmt.Println("Name:", name, "Age:", age, "Height:", height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Height is %.2f\n", height)
	fmt.Printf("Type of name is = %T\n", name)
}
