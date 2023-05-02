package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// delete a file
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a file on the disk
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// write to the file
	file.WriteString(" This is a program written by Nava  which will write to a file !!\n")
	file.WriteString(" second line of the program to write to a file !!")

	// check if string is present in the file
	ok, err := checkStringifExist("test.txt", "Nava")
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		fmt.Println("String found in the file")
	} else {
		fmt.Println("String not found in the file")
	}

	deleteAFile("test.txt")

}

func checkStringifExist(file string, searchPattern string) (ok bool, err error) {
	// open file
	f, err := os.Open(file)
	if err != nil {
		return false, err
	}
	// check if string is present in the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchPattern) {
			return true, nil
		}
	}
	return false, nil
}

func deleteAFile(file string) (status bool, err error) {
	err = os.Remove(file)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

// Create a file
func createAFile(file string) (status bool, err error) {
	// create a file on the disk
	file, err = os.Create(file)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer file.Close()
	return true, nil
}
