package main

import (
	"flag"
	"log"
)

var (
	fFlag int
	dFlag string
	sFlag bool
)

func init() {
	flag.CommandLine.IntVar(&fFlag, "f", 0, "select  only these fields;  also print any line that contains no delimiter character, unless the -s option is specified")
	flag.CommandLine.StringVar(&dFlag, "d", "\t", "specify field delimiter (TAB by default)")
	flag.CommandLine.BoolVar(&sFlag, "s", false, "do not print lines not containing delimiters")
}

func main() {
	flag.Parse()
	if isFlagPassed("f") && fFlag <= 0 {
		log.Fatalln("fields are numbered from 1")
	}

	lines, err := readInput()
	if err != nil {
		log.Fatalln("Couldn't read input.txt")
	}

	for _, line := range lines {
		Cut(line, dFlag)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
