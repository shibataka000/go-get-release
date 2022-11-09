package main

import (
	"fmt"
	"os"

	"github.com/shibataka000/go-get-release/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
