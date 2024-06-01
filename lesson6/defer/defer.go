package main

import "fmt"

func main() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println("work")
	defer fmt.Println("Defer 3")
	fmt.Println("work 2")
}
