package hello

import (
	"golang.org/x/example/hello/reverse"
)

func ReverseHello() string {
	return reverse.String("Hello")
}
