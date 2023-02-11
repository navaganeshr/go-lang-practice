package main

import (
	"fmt"
	"sort"
)

/*
 Arrays in go ?
	* 1. Arrays are fixed length and fixed type
	* 2. Arrays are zero indexed
	* 3. Arrays are value types

	* Initialization of an array:
		var <name> [<length>]<type> = [<length>]<type>{<value>, <value>, <value>, ...}
		* the length of the array is optional
		var marks [5]int = [5]int{41, 52, 33, 64, 95}
		marks := [5]int{41, 52, 33, 64, 95}
		marks := [...]int{41, 52, 33, 64, 95}
*/

func main() {
	fruits := [...]string{"apple", "banana", "orange", "grapes", "mango"}
	marks := [5]int{90, 89, 78, 67, 56}

	//* Display the array
	fmt.Println("marks are : ", marks)
	fmt.Println("number of fruits: ", len(fruits))

	//* sort the array in ascending order
	sort.Strings(fruits[:])
	fmt.Println("number of fruits: ", fruits)

	//* sort the array in descending order
	sort.Strings(fruits[:])
	fmt.Println("number of fruits: ", fruits)

	//* update the value of an array
	fruits[0] = "pineapple"
	fmt.Println("list of  fruits: ", fruits)

	//* iterate over the array
	for i := 0; i < len(fruits); i++ {
		println("fruits at index", i, "is", fruits[i])
		if fruits[i] == "orange" {
			break
		}
	}

	// *range function for array iteration in go
	for index, value := range marks {
		fmt.Println("marks at index", index, "is", value)
	}
}
