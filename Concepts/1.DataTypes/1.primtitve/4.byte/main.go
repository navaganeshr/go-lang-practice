package main

import "fmt"

const (
	Charstart = 65
	Charend   = 90
)

func main() {
	GenerateAlphabets()
}

// Generate English alphabets
func GenerateAlphabets() {
	// %c	the character represented by the corresponding Unicode code point
	// %v the value in a default format
	// %T	a Go-syntax representation of the type of the value

	for i := Charstart; i <= Charend; i++ {
		fmt.Printf("%v \t %c \t %T ", i, i, i)
		fmt.Println()
	}
}
