package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	//message := greetings.Hello("LOBST AND CRABSTER ! ! 🦞🦀")
	fmt.Println(message)
}
