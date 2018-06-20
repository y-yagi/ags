package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	pipeline "github.com/mattn/go-pipeline"
)

var (
	logger *log.Logger
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	var cmds [][]string
	var cmdWithOptions []string
	var patterns []string

	cmdWithOptions = append(cmdWithOptions, "ag")
	for i, arg := range args {
		if isOption(arg) {
			cmdWithOptions = append(cmdWithOptions, arg)
		} else {
			patterns = args[i:]
			break
		}
	}

	if len(patterns) == 0 {
		usage()
		os.Exit(1)
	}

	cmdWithOptions = append(cmdWithOptions, patterns[0])
	cmds = append(cmds, cmdWithOptions)

	for _, pattern := range patterns[1:] {
		cmds = append(cmds, []string{"ag", pattern})
	}

	out, err := pipeline.Output(cmds...)
	if err == nil {
		fmt.Printf(colorized(string(out), patterns))
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] PATTERNS\n", os.Args[0])
}

func isOption(arg string) bool {
	if strings.HasPrefix(arg, "-") {
		return true
	}

	return false
}

func colorized(out string, patterns []string) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	for _, pattern := range patterns {
		out = strings.Replace(out, pattern, yellow(pattern), -1)
	}

	return out
}
