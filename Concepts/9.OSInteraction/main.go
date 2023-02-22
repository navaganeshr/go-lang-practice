package main

import (
	"fmt"
	"os"
)

func main() {
	//take imput argument

	// Read the value from environment variable
	var username = os.Getenv("USERNAME1")
	fmt.Println("Enter passworf for user:", username)

	// shell interactive program
	var password string
	fmt.Scanln(&password)

	go_file_name := os.Args[0]
	fmt.Println("Go file name:", go_file_name)
}
