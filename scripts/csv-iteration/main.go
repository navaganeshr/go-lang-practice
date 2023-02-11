package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	var username string
	var password string
	var option string

	// sign in or sign up
	// if sign in, then ask for username and password
	// if sign up, then ask for username and password and store it in a csv file
	// if username and password are correct, then print "Welcome to the system"
	// if username and password are incorrect, then print "Invalid username or password" and exit the program
	fmt.Println("sign-in or sign-up?")
	fmt.Scanln(&option)
	if option == "sign-in" {
		fmt.Println("Please enter your username")
		fmt.Scanln(&username)
		fmt.Println("Please enter your password")
		fmt.Scanln(&password)
		userSignIn(username, password)
	} else if option == "sign-up" {
		fmt.Println("Please enter your username")
		fmt.Scanln(&username)
		fmt.Println("Please enter your password")
		fmt.Scanln(&password)
		userSignUp(username, password)
	} else {
		fmt.Println("Invalid option")
	}
}

func userSignIn(username string, password string) {
	var is_signed_in bool

	fmt.Println("username is: ", username)
	fmt.Println("password is: ", password)

	//* open the csv file
	recordFile, err := os.Open("credentials.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// *create a new reader
	reader := csv.NewReader(recordFile)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// *iterate through the records
	for _, item := range record {
		if item[0] == username && item[1] == password {
			fmt.Println("Welcome to the system")
			is_signed_in = true
		}
	}

	// *validate users in the file
	if is_signed_in != true {
		fmt.Println("Invalid username or password", is_signed_in)
	}
}

func userSignUp(username string, password string) {
	fmt.Println("username is: ", username)
	fmt.Println("password is: ", password)

	//* check if the username already exists
	//* open the csv file
	recordFile, err := os.Open("credentials.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := csv.NewReader(recordFile)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range record {
		if item[0] == username {
			fmt.Println("Username already exists")
			return
		}
	}

	// *create a new writer
	writer := csv.NewWriter(recordFile)
	var user = []string{username, password}
	writer.WriteAll([][]string{user})
}
