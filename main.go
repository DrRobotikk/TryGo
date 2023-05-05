package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	messege := greetings.Hello("Robin")
	fmt.Println(messege)
}