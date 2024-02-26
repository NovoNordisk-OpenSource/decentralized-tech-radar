// Pretty sure this is done so that you can have the package hosted on pkg.go.dev so that others go get
// But please do keep in mind when we import this in the main.go file we are actually importing from the local
// Module aka this file. It's very confusing but just know import "github...." refers to the local mod file
// and the go get stuff is for the hosted version of this pkg....
module github.com/Agile-Arch-Angels/decentralized-tech-radar_dev

go 1.21.0
