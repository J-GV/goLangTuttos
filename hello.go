package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("MOTHERFUCKER ! !")
	fmt.Println(message)
}
