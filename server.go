package main

import (
	"flag"
	"fmt"

	r "github.com/tavvfiq/cafe-rest-api/routers"
)

func main() {
	handleArgs()
}

// handle input arguments
func handleArgs() {
	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case "run":
		r.Start()
	case "seed":
		if len(args[1:]) == 0 {
			fmt.Println("Seeding All")
		} else {
			fmt.Printf("Seeding %v", args[1:])
		}
	default:
		r.Start()
	}

}
