package main

import "fmt"

type abstracttype interface {
	int64 | float64 | string
}

func main() {
	int_slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	string_slice := []string{"Hello ", "my ", "friend!"}
	fmt.Println(AddGeneric(int_slice))
	fmt.Println(AddGeneric(string_slice))
}

func AddGeneric[T abstracttype](s []T) T {
	var result T
	for _, value := range s {
		result = result + value
	}
	return result
}
