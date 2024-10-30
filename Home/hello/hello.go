package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("LORD AND MASTER ! !")
	fmt.Println(message)
}
