package main

import "fmt"

func SumAndSubtract(x, y float32) (sum, difference float32) {
	sum = x + y
	difference = x - y
	return
}

func main() {
	sum, difference := SumAndSubtract(5.4, 3.7)
	fmt.Println("Sum and sub results:", sum, difference) 
}
