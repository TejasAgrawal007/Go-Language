package main

import (
	"fmt"
	myutils "mylerning/myUtils"
)

func main() {
	fmt.Println("Hey Go Lang I am Tejas")
	myutils.PrintMessage("Hello Tejas I am Utils from myUtils")

	var name string = "Tejas Agrawal Here"
	var version = "lastest version"
	var diminsion float64 = 87.35
	fmt.Println(name)
	fmt.Println(version)
	fmt.Println("Diminsion = ", diminsion)
}
