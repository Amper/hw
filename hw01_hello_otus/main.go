package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	hello := "Hello, OTUS!"
	reversedHello := stringutil.Reverse(hello)

	fmt.Print(reversedHello)
}
