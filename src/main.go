package main

import(
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	fmt.Printf("Hello, %s\n", *name)
}