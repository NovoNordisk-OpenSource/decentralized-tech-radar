package HTML

import (
	"embed"
	"html/template"
	"log"
	"os"
)

//go:embed css/style.css
var styleCSS embed.FS

var htmlFileName string = "index"

func GenerateHtml(csvData string) {
	const tmpl = `
	<!doctype html>
	<html lang="en">

	<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<link href="../src/HTML/images/favicon.ico" rel="icon" />
	<style>
		{{.CSS}}
	</style>
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=Noto+Sans:ital,wght@0,100..900;1,100..900&display=swap" rel="stylesheet">
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
				<p class="agree-terms">Visit the Novo Nordisk <a href="https://github.com/NovoNordisk-OpenSource/decentralized-tech-radar">Tech Radar open source on Github</a>. This tech radar was heavily inspired by <a href="https://github.com/thoughtworks/build-your-own-radar/">Thoughtworks</a>.</p>
			</div>
		</main>
	</body>
	<script src="./HTML/js/libraries/require.js"></script>
	<script src="./HTML/js/requireConfig.js"></script>

	<!-- this script builds the radar with the go generated csv file -->
	<script>
		require(['./HTML/js/renderingTechRadar.js'], function(Factory) {
			Factory({{.CSV}}).build(); //{{.}} refers to the csvData
		})
	</script>
	</html>
	`

	// Read the content of the CSS file
	cssContent, err := styleCSS.ReadFile("css/style.css")
	if err != nil {
		log.Fatalf("failed to read embedded CSS file: %v", err)
	}

	// Create a template data structure to hold both CSS and CSV data
	data := struct {
		CSS template.CSS
		CSV string
	}{
		CSS: template.CSS(cssContent),
		CSV: csvData,
	}

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
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
