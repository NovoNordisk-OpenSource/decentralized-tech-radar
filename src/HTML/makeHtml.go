package HTML

import (
	Reader "github.com/NovoNordisk-OpenSource/decentralized-tech-radar/SpecReader"
	"html/template"
	"log"
	"os"
)

var htmlFileName string = "index"

func GenerateHtml(blips Reader.Blips) {
	const tmpl = `
	<!doctype html>
	<html lang="en">

	<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<link href="/images/favicon.ico" rel="icon" />
	<link rel="preconnect" href="https://rsms.me/" />
	<link rel="stylesheet" href="https://rsms.me/inter/inter.css" integrity="sha512-byor" />
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link href="https://fonts.googleapis.com/css2?family=Bitter:wght@700&display=swap" rel="stylesheet" />
	
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-autocomplete/1.0.7/jquery.auto-complete.min.js"></script>
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
		<div class="graph-footer"></div>
	</main>

	</body>

	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-autocomplete/1.0.7/jquery.auto-complete.min.js"></script>
	<script src="https://d3js.org/d3.v7.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/lodash@4.17.21/lodash.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.3.6/require.min.js"></script>
	<script src="./js/requireConfig.js"></script>
	<script>
		require(['./js/site.js'], function() {
		
		});
  	</script>



	<!--<script src="./models/radar.js"></script>
	<script src="./models/quadrant-js"></script>
	<script src="./models/ring.js"></script>
	<script src="./models/blip.js"></script>
	<script src="./graphing/radar.js"></script>
	<script src="./js/config.js"></script>
	<script>
	const featureToggles = config().featureToggles
	</script>
	<script src="./js/graphing/config.js"></script>
	
	<script src="./js/util/factory.js"></script>
	<script src="./js/site.js"></script>-->
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
