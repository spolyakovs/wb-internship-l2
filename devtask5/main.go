package main

import (
	"flag"
	"log"
	"runtime"

	myLog "github.com/spolyakovs/wb-internship-l2/devtask5/log"
	"github.com/spolyakovs/wb-internship-l2/devtask5/pattern"
)

var (
	redBoldColor string
	resetColor   string

	aCapFlag int
	bCapFlag int
	cCapFlag int
	cFlag    bool
	iFlag    bool
	vFlag    bool
	fCapFlag bool
	nFlag    bool
)

func init() {
	if runtime.GOOS == "windows" {
		redBoldColor = "" // windows console doesn't support colors
		resetColor = ""
	} else {
		redBoldColor = "\033[1m\033[31m"
		resetColor = "\033[0m"
	}

	flag.CommandLine.IntVar(&cCapFlag, "C", 0, "lines to print after and before matching lines")
	// basically equal -A=n -B=n
	// if -A or -B flags specified, they take priority over -C
	flag.CommandLine.IntVar(&aCapFlag, "A", 0, "lines to print after matching lines")
	flag.CommandLine.IntVar(&bCapFlag, "B", 0, "lines to print before matching lines")

	flag.CommandLine.BoolVar(&cFlag, "c", false, "print number of matched lines in each file instead of lines themselves")
	flag.CommandLine.BoolVar(&iFlag, "i", false, "ignore case (with this flag patterns that only differ in case will match)")
	flag.CommandLine.BoolVar(&vFlag, "v", false, "inverse matching (match lines that DO NOT match pattern)")
	flag.CommandLine.BoolVar(&fCapFlag, "F", false, "matches line ONLY if pattern matches whole line")
	flag.CommandLine.BoolVar(&nFlag, "n", false, "also prints line number in file")
}

func main() {
	flag.Parse()
	if isFlagPassed("C") {
		if !isFlagPassed("A") {
			aCapFlag = cCapFlag
		}

		if !isFlagPassed("B") {
			bCapFlag = cCapFlag
		}
	}

	patternStr := flag.Arg(0)
	filenames := flag.Args()[1:]
	needFilename := len(filenames) > 1 // TODO comment

	for _, file := range filenames {
		lines, err := readFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		if !needFilename {
			file = ""
		}

		var logger myLog.Logger
		if cFlag {
			logger = myLog.NewLoggerCount(file)
		} else {
			logger = myLog.NewLoggerDefault(file, nFlag)
		}

		var matcher pattern.Matcher
		if fCapFlag {
			matcher = pattern.NewPatternLine(patternStr)
		} else {
			matcher = pattern.NewPatternDefault(patternStr)
		}

		// TODO
		Match(patternStr, lines, matcher, logger)
		// sortLines(lines)
		// previousLine := ""
		// for _, line := range lines {

		// 	if !uFlag || previousLine != line {
		// 		output.WriteString(line + "\n")
		// 	}
		// 	previousLine = line
		// }
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
