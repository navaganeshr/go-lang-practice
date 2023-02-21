package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Random number generator")
	regset := math.Pow(10, 10)
	fmt.Println("Random number between 0 and", regset)
	fmt.Println(randomUniqueNumber(int(regset)))
}

// six digit random number
func randomUniqueNumber(randlen int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(randlen)
	fmt.Println("Random number is", result)
	// if numbers are repeating, call the function again
	if validateRandomNumber(result) {
		return randomUniqueNumber(randlen)
	}
	return
}

func validateRandomNumber(number int) bool {
	num_str := strconv.Itoa(number)
	for i := 0; i < len(num_str); i++ {
		for j := i + 1; j < len(num_str); j++ {
			fmt.Println("num_str:", num_str, "origin:", num_str[i], "Dummy", num_str[j])
			if num_str[i] == num_str[j] {
				return false
			}
		}
	}
	return true
}
