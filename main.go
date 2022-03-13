package main

import (
	"fmt"
	"github.com/johannesscr/go-playground/example"
	"log"
)

// main is where all the examples can be executed from.
func main() {
	var o string
	fmt.Print("Go Playground\n\n")
	fmt.Print("Please choose the example you would like to run:\n")
	fmt.Print("a: read user input from the console/command line.\n")
	fmt.Print("> : ")
	_, err := fmt.Scanln(&o)
	if err != nil {
		log.Fatalln(err)
	}
	switch o {
	case "a":
		example.UserInput()
		break
	}
}
