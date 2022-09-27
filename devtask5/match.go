package main

import (
	"strings"

	logger "github.com/spolyakovs/wb-internship-l2/devtask5/log"
	"github.com/spolyakovs/wb-internship-l2/devtask5/pattern"
)

// TODO: new loggers
func Match(p string, lines []string, matcher pattern.Matcher, logger logger.Logger) {
	var lastAddedLineNum int = 0
	afterQueue := make(map[int]bool, aCapFlag)

	for i, line := range lines {

		pMatches := matcher.Match(line)
		found := len(pMatches) > 0

		if found {
			line = matchedString(line, pMatches)
			lines[i] = line
		}

		matched := (found != vFlag) // matched reverses "found" if vFlag set

		if !matched {
			if afterQueue[i] {
				logger.Put(i+1, line, matched)
				lastAddedLineNum = i + 1
				delete(afterQueue, i)
			}
			continue
		}

		// adds "--" if there is gap between lines
		if (i+1)-bCapFlag-lastAddedLineNum > 1 && lastAddedLineNum != 0 {
			logger.Put(0, "--", false)
		}

		for bIndex := i - bCapFlag; bIndex < i; bIndex++ {

			if bIndex+1 <= lastAddedLineNum {
				continue
			}

			delete(afterQueue, bIndex)

			logger.Put(bIndex+1, lines[bIndex], false)
			lastAddedLineNum = bIndex + 1
		}

		logger.Put(i+1, line, matched)
		lastAddedLineNum = i + 1
		delete(afterQueue, i)

		// add line afterQueue according to aCapFlag
		for aIndex := i + 1; aIndex <= i+aCapFlag; aIndex++ {
			if aIndex >= len(lines) {
				break
			}

			afterQueue[aIndex] = true
		}
	}

	logger.Log()
}

// matchedString returns line with matched pattern highlighted in red
func matchedString(line string, pMatches []pattern.PatternMatch) string {
	if len(pMatches) == 0 {
		return line
	}

	var builder strings.Builder
	builder.Grow(len(line) + len(pMatches)*(len(redBoldColor)+len(resetColor)))

	lineRunes := []rune(line)
	lastWrittenRuneIndex := -1

	for _, pMatch := range pMatches {
		for _, r := range lineRunes[lastWrittenRuneIndex+1 : pMatch.Begin] {
			builder.WriteRune(r)
			lastWrittenRuneIndex++
		}

		builder.WriteString(redBoldColor)

		for _, r := range lineRunes[pMatch.Begin:pMatch.End] {
			builder.WriteRune(r)
			lastWrittenRuneIndex++
		}

		builder.WriteString(resetColor)
	}

	for _, r := range lineRunes[pMatches[len(pMatches)-1].End:] {
		builder.WriteRune(r)
	}

	return builder.String()

}
