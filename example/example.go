package main

import (
	"config-parser"
	"fmt"
	"log"
	"os"
)

func main() {
	source, err := os.ReadFile("./example/config")
	if err != nil {
		log.Fatal(err)
	}

	config := cfg_parser.NewConfig(string(source))
	config.Parse()

	result, err := config.GrabPairs("Information")
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range result {
		fmt.Printf("key = %s | value = %s\n", key, value)
	}

	// method 1
	author := result["Author"]
	fmt.Println("Author for Information is", author)

	// method 2
	_, value, err := config.GrabPair("Version-0.2", "Author")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Author for version 2 is", value)
}
