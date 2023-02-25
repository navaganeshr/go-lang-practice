package main

import "fmt"

func main() {
	// create user struct
	type User struct {
		FirstName string
		LastName  string
		Age       int
	}
	var nava User
	nava.FirstName = "Nava"
	nava.LastName = "Ganesh"
	nava.Age = 21

	John := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       21,
	}

	Andy := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       21,
	}

	// print struct
	fmt.Printf("Nava's Details: %+v \n", nava)
	fmt.Printf("John's Details: %+v \n", John)
	fmt.Printf("Andy's Details: %+v \n", Andy)
	if compareStructs(Andy, John) {
		fmt.Println("Andy and John are the same person")
	} else {
		fmt.Println("Andy and John are not the same person")
	}
}

func compareStructs(struct1 struct {
	FirstName string
	LastName  string
	Age       int
}, struct2 struct {
	FirstName string
	LastName  string
	Age       int
}) bool {
	return struct1 == struct2
}
