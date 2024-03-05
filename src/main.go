package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/src/SayHello"
)

func main() {
	name := flag.String("name", "world", "This takes a name/string ")
	flag.Parse()

	greeting := SayHello.SayHello(*name)

	makeHtml(greeting)
}

type TechList struct {
	Title string
	Adopt bool
	Drop  bool
	Trial bool
}

type PageData struct {
	PageTitle string     // page title
	Greeting  string     // greeting placeholder text
	TechList  []TechList // list of items
}

// https://www.educative.io/answers/what-is-html-template-in-golang
func makeHtml(htmldata string) {
	const tmpl = `
	<html>
		<head>
			<title>{{.PageTitle}}</title>
			<link rel="stylesheet" href="css/style.css" type="text/css">
		</head>
		<body>
			<h1 class="pageTitle">{{.PageTitle}}</h1>
			<span>{{.Greeting}}</span>
			<ul>
				{{range .TechList}}
					{{if .Adopt}}
						<li>{{.Title}} &#10004</li>
					{{else}}
						<li>{{.Title}}</li>
					{{end}}
				{{end}}
			</ul>
			<a href="//github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/" title="Tech radar" ref="rel">Tech radar repo</a>
		</body>
	</html>
	`

	// Make and parse the HTML template
	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Initialze a struct storing page data and todo data
	data := PageData{
		//show a string of data
		PageTitle: "Title for page",
		//show a string of data
		Greeting: htmldata,
		//show a list of data
		TechList: []TechList{
			{Title: "C++", Adopt: false},
			{Title: "Java", Adopt: true},
			{Title: "F#", Adopt: true},
			{Title: "Golang", Adopt: true},
		},
	}

	//remove the index file if there's an old one
	if err := os.Remove("index.html"); err != nil {
		log.Fatal(err)
	}

	// Open index.html for writing (create if it doesn't exist)
	file, err := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0644) // 0644 grants the owner read and write access
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//execute the html and data
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
