package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Bird struct {
		Species     string
		Description string
	}
	var bird Bird
	birdJson := ``
	fmt.Printf("birdJson type: %T \n", birdJson)
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	file, err := os.Open("bird.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// parse json objects from a file
// https://stackoverflow.com/questions/21362931/how-to-parse-json-objects-from-a-file-in-go
