package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 42
	a := "10"

	if n, err := strconv.Atoi(a); err == nil {
		fmt.Println("n is positive", n)
	} else {
		fmt.Println("n is negative", n)
	}
}
