package pattern

type PatternMatch struct {
	Begin int // rune(!!!) number in a line for beginning of matched pattern
	End   int // rune(!!!) number in a line for ending of matched pattern
	// string[end] ISN'T in the matched pattern
}

type Matcher interface {
	Match(string) []PatternMatch
}
