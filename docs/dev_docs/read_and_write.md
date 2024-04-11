# Preface
*WIP this document will be updated throughout the project*

# Reading
Technologies for the radar (or blips) are stored in the program using structs. These structs can then be passed between different parts of the program.

## Reading Spec files from CSV format
For reading spec files in csv format, the `ReadCsvSpec` function is used. Since we know the structure of our data we can predefine our structs to fit that structure.  
By using the [gocarina/gocsv](https://github.com/gocarina/gocsv) module we can  unmarshal the csv file. Taking the csv file and turning it into our data structure.

# Writing
**Will be rewritten**  
This opens a index.html file, and if an index.html file doesn't exists, it creates it. It also removes any old index.html files if there are any.  

## Generating html
This feature is used when a user already has a csv specfile on they device and wishes to use that for generating html for the tech-radar.

Generating is done by:
```go
func GenerateHtml(blips Reader.Blips)
```
This takes a struct of blips, which was made from the given csv file, that then get translated into html with GoLang template.

It uses a GoLang template, which is a data-driven template for generating html,
which is safe against code injection and/or XSS attacks.
GoLang's own site on this package: https://pkg.go.dev/html/template

Currently it can convert data from a CLI command to html.
