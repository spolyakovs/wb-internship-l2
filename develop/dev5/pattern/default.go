package pattern

type patternDefault struct {
	pattern string
}

func NewPatternDefault(pattern string) *patternDefault {
	return &patternDefault{
		pattern: pattern,
	}
}

func (p patternDefault) Match(line string) (pMatches []PatternMatch) {
	if len(p.pattern) > len(line) {
		return pMatches
	}
	patternRunes := []rune(p.pattern)
	lineRunes := []rune(line)

	// iterate over line runes that can be beginning of pattern
	for i := range lineRunes[:len(lineRunes)-len(patternRunes)+1] {
		// compare pattern and current part of line with equal length
		if string(lineRunes[i:i+len(patternRunes)]) == string(patternRunes) {
			pMatch := PatternMatch{
				Begin: i,
				End:   i + len(patternRunes),
			}
			if len(pMatches) == 0 || pMatches[len(pMatches)-1].End <= pMatch.Begin {
				pMatches = append(pMatches, pMatch)
			}
		}
	}

	return pMatches
}
