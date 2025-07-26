// cmd/datarepository/main.go
package main

import (
	"flag"
	"log"
	"os"

	"datarepository/internal/datarepository"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := datarepository.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
