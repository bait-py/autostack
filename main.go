package main

import (
	"log"

	"github.com/bait-py/autostack/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
