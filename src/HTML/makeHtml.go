package HTML

import (
	"html/template"
	"log"
	"os"
)

var htmlFileName string = "index"

func GenerateHtml(csvData string) {
	const tmpl = `
	<!doctype html>
	<html lang="en">

	<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<link href="./js/images/favicon.ico" rel="icon" />
	<link rel="preconnect" href="https://rsms.me/" />
	<link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link href="./js/stylesheets/style.css" rel="stylesheet"/>
	<link href="https://fonts.googleapis.com/css2?family=Bitter:wght@700&display=swap" rel="stylesheet" />
	</head>

	<body>
		<main>
			<div class="input-sheet-form home-page">
			<p>
				building the radar...
			</p>
			</div>
			<div class="graph-header"></div>
			<div id="radar">
			<p class="no-blip-text">
				There are no blips on this quadrant, please check your CSV file.
			</p>
			</div>
			<div class="all-quadrants-mobile show-all-quadrants-mobile"></div>
			<div class="graph-footer">
				<p class="agree-terms">Visit the Novo Nordisk <a href="https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar">Tech Radar open source on Github</a>. This tech radar was heavily inspired by Thoughtworks.</p>
			</div>
		</main>
	</body>

	<script src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.3.6/require.min.js"></script>
	<script src="./js/requireConfig.js"></script>

	<!-- this script builds the radar with the go generated csv file -->
	<script>
		require(['./js/util/factory.js'], function(Factory) {
			Factory({{.}}).build(); //{{.}} refers to the csvData
		})
	</script>
	</html>
	`
	// Make and parse the HTML template
	t, err := template.New(htmlFileName).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Open index.html for writing (create if it doesn't exist)
	os.Remove(htmlFileName + ".html")
	file, err := os.OpenFile(htmlFileName+".html", os.O_WRONLY|os.O_CREATE, 0644) // 0644 grants the owner read and write access
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//execute the html and data
	err = t.Execute(file, csvData)
	if err != nil {
		panic(err)
	}
}
