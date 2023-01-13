package main

import (
	"log"

	"github.com/shibataka000/go-get-release/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
