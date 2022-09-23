package main

import (
	"flag"
	"log"
	"os"
)

var (
	kFlag int
	nFlag bool
	rFlag bool
	uFlag bool
	// mFlag bool
	// bFlag bool
	// cFlag bool
	// hFlag bool
)

func init() {
	flag.CommandLine.IntVar(&kFlag, "k", 0, "column number to sort by")
	flag.CommandLine.BoolVar(&nFlag, "n", false, "sort by numeric value")
	flag.CommandLine.BoolVar(&rFlag, "r", false, "reverse sorting order")
	flag.CommandLine.BoolVar(&uFlag, "u", false, "do not output repeat lines")
}

func main() {
	flag.Parse()
	filenames := flag.Args()

	for _, file := range filenames {
		lines, err := readFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		output, err := os.Create("output_" + file)
		if err != nil {
			log.Fatalln(err)
		}
		defer output.Close()

		// TODO: cFlag
		sortLines(lines)
		previousLine := ""
		for _, line := range lines {

			if !uFlag || previousLine != line {
				output.WriteString(line + "\n")
			}
			previousLine = line
		}
		// TODO: >> output.txt
	}
}
