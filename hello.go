package main

import (
	"fmt"
	"peyton-spencer/gin-helloworld/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, go")

	// local package
	fmt.Println(morestrings.ReverseRunes("!oooG ,olleH"))

	// remote package
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
