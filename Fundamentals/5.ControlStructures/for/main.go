package main

import "fmt"

func main() {

TestLabel:
	for x := 0; x < 20; x++ {
		if x == 15 {
			break TestLabel
		}
		fmt.Printf("value of x: %d\n", x)
	}
}
