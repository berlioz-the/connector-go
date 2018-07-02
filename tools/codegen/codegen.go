package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CWD: %s\n", dir)

	generateAws()
}
