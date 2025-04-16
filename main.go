package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello", "World", 123, true)
	fmt.Println(quote.Opt())
}
