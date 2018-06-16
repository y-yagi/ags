package main

import (
	"fmt"
	"log"
	"os"

	pipeline "github.com/mattn/go-pipeline"
)

var (
	logger *log.Logger
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s PATTERNS\n", os.Args[0])
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	var cmds [][]string

	for _, pattern := range args {
		cmds = append(cmds, []string{"ag", pattern})
	}

	out, err := pipeline.Output(cmds...)
	if err == nil {
		fmt.Printf(string(out))
	}
}
