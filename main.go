package main

import (
	"flag"
	"fmt"
)

func main() {
	test := flag.String("t", "test", "test flag")
	flag.Parse()

	fmt.Printf("simple or not? %q\n", *test)
}
