package main

import "fmt"

// func divide(a, b float64) (float64, error) {

// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator must not be zero!")
// 	}
// 	return a / b, nil
// }

func divide(a, b float64) (float64, string) {

	if b == 0 {
		return 0, "denominator must not be zero!"
	}
	return a / b, "nil"
}
func main() {
	fmt.Println("I am Working")
	ans, _ := divide(10, 4)
	fmt.Println("The Divison of Two Num is ", ans)
}
