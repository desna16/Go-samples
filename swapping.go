package main

import "fmt"

func main() {
	var first int = 3
	fmt.Print("first number : \n", first)

	var second int = 1
	fmt.Print("second number : \n ", second)

	first = first - second
	second = first + second
	first = second - first

	fmt.Println("First number : \n", first)
	fmt.Println("Second number :  ", second)
}