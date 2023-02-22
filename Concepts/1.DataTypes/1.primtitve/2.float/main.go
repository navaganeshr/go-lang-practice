package main

import "fmt"

func main() {
	var float_1 float64 = 1.123
	var float_2 float64
	float_3 := 10.90
	float_2 = 100.345678
	fmt.Printf("integer_1: %v \n integer_2 %v \n integer_3 %v \n", float_1, float_2, float_3)
	fmt.Println("Sum of 3 numbers: ", sum(float_1, float_2, float_3))
}

// sum of 3 numbers
func sum(a float64, b float64, c float64) float64 {
	return a + b + c
}
