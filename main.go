package main

import (
	"log"

	"github.com/shibataka000/go-get-release/cmd"
)

func main() {
	log.SetFlags(0)
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
