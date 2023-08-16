package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("todo.json")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully opened JSON")
	}
	defer jsonFile.Close()
}
