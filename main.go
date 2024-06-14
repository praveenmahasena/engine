package main

import (
	"log"

	"github.com/praveenmahasena/underwriteengine/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
