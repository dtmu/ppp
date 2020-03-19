package main

import (
	"log"
	"os"
)

// Entry point
func main() {
	err := cmd()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	setHandler()

	if err = startWebServer(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	os.Exit(0)
}
