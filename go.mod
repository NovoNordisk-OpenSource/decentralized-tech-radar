// Pretty sure this is done so that you can have the package hosted on pkg.go.dev so that others go get
// But please do keep in mind when we import this in the main.go file we are actually importing from the local
// Module aka this file. It's very confusing but just know import "github...." refers to the local mod file
// and the go get stuff is for the hosted version of this pkg....
module github.com/Agile-Arch-Angels/decentralized-tech-radar_dev

go 1.21.0

require (
	dagger.io/dagger v0.10.0
	github.com/gocarina/gocsv v0.0.0-20231116093920-b87c2d0e983a
)

require (
	github.com/99designs/gqlgen v0.17.31 // indirect
	github.com/Khan/genqlient v0.6.0 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/vektah/gqlparser/v2 v2.5.6 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
