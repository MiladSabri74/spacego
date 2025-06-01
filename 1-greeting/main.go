package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "greeting")

	flag.Parse()

	fmt.Printf("Hello ,%s!\n", *name)
}
