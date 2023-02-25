package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// check log file eixsts
	LOG_FILE := os.Getenv("LOG_FILE")

	fmt.Println("INFO: Application started !! ")
	fmt.Printf("INFO: Parsing log file %v \n", LOG_FILE)
	if _, err := os.Stat(LOG_FILE); os.IsNotExist(err) {
		fmt.Printf("ERROR: Log file %v is not present", LOG_FILE)
		os.Exit(1)
	}
	// Read log file
	readLogFile(LOG_FILE)

}

func readLogFile(LOG_FILE string) {
	// var log map[string]interface{}
	// open log file
	// file, err := os.Open(LOG_FILE)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// read log file line by line
	scanner := bufio.NewScanner(os.Stdin)
	var input []byte
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		input = append(input, scanner.Bytes()...)
		// json.Unmarshal([]byte(scanner.Text()), &log)
	}
	fmt.Println(string(input))

}
