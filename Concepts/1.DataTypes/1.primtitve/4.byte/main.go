package main

import "fmt"

func main() {
	Charstart := 65
	Charend := 90
	var someByte byte = 'A'
	fmt.Printf("value: %v \n byte: %c \t Type: %T \n", someByte, someByte, someByte)
	GenerateAlphabets(Charstart, Charend)
}

// Generate English alphabets
func GenerateAlphabets(charstart int, charend int) {
	// %c	the character represented by the corresponding Unicode code point
	// %v the value in a default format
	// %T	a Go-syntax representation of the type of the value
	for i := charstart; i <= charend; i++ {
		fmt.Printf("%v \t %c \t %T ", i, i, i)
		fmt.Println()
	}

}
