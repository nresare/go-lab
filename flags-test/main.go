package main

import (
	"fmt"
	"github.com/thought-machine/go-flags"
	"os"
)

type Options struct {
	Version bool `short:"v" long:"version" description:"Print the version"`
	//Targets []string `positional-args:"yes" required:"yes" description:"The targets to build"`
	Args struct{ Targets []string } `positional-args:"yes" required:"yes" description:"All relevant targets"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	fmt.Printf("Doing some work here: targets is '%s'\n", options.Args.Targets)
}
