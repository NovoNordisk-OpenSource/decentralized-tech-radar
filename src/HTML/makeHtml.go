package HTML

import (
	Reader "github.com/Agile-Arch-Angels/decentralized-tech-radar_dev/SpecReader"
	"html/template"
	"log"
	"os"
)

var htmlFileName string = "index"

func GenerateHtml(blips Reader.Blips) {
	const tmpl = `<html>
	<head>
		<title>Header 1</title>
	</head>
	<body>
		<h1 class="pageTitle">Header 1</h1>
		<ul>
			{{range .Blips}}
					<li>Name: {{.Name}}</li>
					<li>Quadrant: {{.Quadrant}}</li>
					<li>Ring: {{.Ring}}</li>
					<li>Is new: {{.IsNew}}</li>
					<li>Moved: {{.Moved}}</li>
					<li>Desc: {{.Description}}</li>
			{{end}}
		</ul>
	</body>
</html>
	`

	// Make and parse the HTML template
	t, err := template.New(htmlFileName).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Open index.html for writing (create if it doesn't exist)
	os.Remove(htmlFileName + ".html")
	file, err := os.OpenFile(htmlFileName + ".html", os.O_WRONLY|os.O_CREATE, 0644) // 0644 grants the owner read and write access
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//execute the html and data
	err = t.Execute(file, blips)
	if err != nil {
		panic(err)
	}
}
