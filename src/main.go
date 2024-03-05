package main

import (
	"flag"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
	view "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/View"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	greeting := SayHello.SayHello(*name)

	view.MakeHtml(greeting)
}
