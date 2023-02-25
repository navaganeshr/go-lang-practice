package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
		x := 20
		fmt.Println("x is: ", x)
	} else {
		fmt.Println("x is less than 5")
	}

	fmt.Println("x is: ", x)

}
