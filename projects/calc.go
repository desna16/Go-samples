package main

import "fmt"

func main() {
	var operation string = "mul"
	var A int = 90
	var B int = 10
	var result int

	switch operation {
	case "add":
		result = (A + B)
	case "sub":
		result = (A - B)
	case "mul":
		result = (A * B)
	case "div":
		result = (A / B)
	default:
		result = 0
	}
	switch {
	case "add" == string(A+B):
		fmt.Printf("sum is %v = ", result)
	case "sub" == string(A-B):
		fmt.Printf("difference is %v = ", result)
	case "mul" == string(A*B):
		fmt.Printf("product is %v = ", result)
	case "div" == string(A/B):
		fmt.Printf("qoutient is %v = ", result)
	default:
		fmt.Printf("Invalid operation\n")
	}
	fmt.Printf("Your result is  %v\n", result)
}
