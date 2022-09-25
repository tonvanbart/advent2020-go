package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("nr of rules %d\n", len(input))

}
