package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	greeting := SayHello.SayHello(*name)

	// Open index.html for writing (create if it doesn't exist)
	file, err := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0644) // 0644 grants the owner read and write access
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the greeting as plain text to the file
	_, err = file.WriteString(greeting + "\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("Wrote to index.html")
}
