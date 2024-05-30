package main

import (
	"fmt"

	"example.com/goodbye"
	"example.com/greetings"
	"example.com/hello"
)

func main() {
	greetings := greetings.DisplayHello()
	hello_reverse := hello.ReverseHello()
	goodbye := goodbye.SayGoodbye()
	fmt.Printf("%v\n", greetings)
	fmt.Printf("%v\n", hello_reverse)
	fmt.Printf("%v\n", goodbye)
}
