package main

import (
	"log"

	"github.com/praveenmahasena/engine/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		log.Fatalln(err)
	}
}
