// Pretty sure this is done so that you can have the package hosted on pkg.go.dev so that others go get
// But please do keep in mind when we import this in the main.go file we are actually importing from the local
// Module aka this file. It's very confusing but just know import "github...." refers to the local mod file
// and the go get stuff is for the hosted version of this pkg....
module github.com/NovoNordisk-OpenSource/decentralized-tech-radar

go 1.21.0

require (
	dagger.io/dagger v0.10.3
	github.com/gocarina/gocsv v0.0.0-20231116093920-b87c2d0e983a
)

require (
	github.com/99designs/gqlgen v0.17.45 // indirect
	github.com/Khan/genqlient v0.7.0 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alexflint/go-arg v1.4.2 // indirect
	github.com/alexflint/go-scalar v1.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.2.0 // indirect
	github.com/urfave/cli/v2 v2.27.1 // indirect
	github.com/vektah/gqlparser/v2 v2.5.11 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
