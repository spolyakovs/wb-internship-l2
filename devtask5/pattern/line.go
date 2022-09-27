package pattern

type patternLine struct {
	pattern string
}

func NewPatternLine(pattern string) *patternLine {
	return &patternLine{
		pattern: pattern,
	}
}

// patternLine.Match matches pattern only to whole line
func (p patternLine) Match(line string) (pMatches []PatternMatch) {
	if p.pattern == line {
		pMatch := PatternMatch{
			Begin: 0,
			End:   len([]rune(line)),
		}
		pMatches = append(pMatches, pMatch)
		return pMatches
	}

	return pMatches
}
