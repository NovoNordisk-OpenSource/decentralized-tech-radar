package main

import(
	"flag"
	"fmt"
	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	fmt.Println(SayHello.SayHello(*name))
}