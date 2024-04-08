package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey what's your Name: ")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Have a Good Day", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Have a Good Day", name)
}
