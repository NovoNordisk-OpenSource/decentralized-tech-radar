# Preface
*WIP this document will be updated throughout the project*

# Reading
Example:
```

```


# Writing
**Will be rewritten**  
This opens a index.html file, and if an index.html file doesn't exists, it creates it. It also removes any old index.html files if there are any.  

It uses a GoLang template, which is a data-driven template for generating html,
which is safe against code injection and/or XSS attacks.
GoLang's own site on this package: https://pkg.go.dev/html/template

Currently it can convert data from a CLI command to html.


Example:
```
    const tmpl = `<span>{{.Greeting}}</span>'

    t, err := template.New("index").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Initialze a struct storing page data and todo data
	data := PageData{
		Greeting: htmldata,
	}

	//remove the index file if there's an old one
    ...

	// Open index.html for writing (create if it doesn't exist)
    ...

	//execute the html and data
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
```

