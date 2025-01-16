package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	processor "github.com/mcastellin/focalfinder/pkg/cmd"
)

type arrayStringFlags []string

func (a *arrayStringFlags) String() string {
	return strings.Join(*a, ", ")
}

func (a *arrayStringFlags) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func main() {
	// Parse CLI arguments
	var factors arrayStringFlags
	flag.Var(&factors, "factor", "Path to the directory containing images")
	flag.Parse()

	dirs := flag.Args()
	if len(dirs) == 0 {
		fmt.Println("Usage: focalfinder [ --factor=\"Make,Model,Factor\" ...] <path_to_directory ...>")
		os.Exit(1)
	}

	processor.ProcessImages(dirs, factors)
}
