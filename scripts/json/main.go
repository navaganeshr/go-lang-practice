package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	f, err := os.ReadFile("test.json")
	if err != nil {
		log.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(f), &data)

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}

}
