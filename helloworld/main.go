package main

//very first package must declare the package that it belongs to
//if 3 files all belong to same project, have to declare `package maing` at the top

import "fmt"

//used to import standard library package available to the go language
//fmt library used to print out information by the terminal
//can import in packages authored by other engineers

func main() {
	fmt.Println("Hi, There!")
}

//print L N, println

//TODO: how do we run our code in the project?
//go run -your file.go-

//TODO: go cli commands
// go build -> compiles go source code files
// go run -> compiles AND executes, one or two files
// go fmt -> formats all the code in each file in the current directory
// go install -> compiles and installs a package
// go get -> compiles the raw source code of someone else's package
// go test -> run any tests associated with the current project

//TODO:what does 'package main' mean?
// package = project = workspace
// 2 types of packages: EXECUTABLE VS REUSABLE
// executable: generates a file that can be run
// reusable: code used as `helpers`, good place for reusable logic
// name of package, [package main], is what determines wether you're making ex or re type package
// only use package main for files that want to be run
// any other name is for a reusable package, main packages are for exe type packages

//TODO: what does 'import fmt' mean?

//TODO: what is a func?

//TODO: how is the 'main.go' file organized?
