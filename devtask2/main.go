package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrIncorrectString = errors.New("incorrect string")

// pair of rune and count to repeat this rune
// token.count - TODO
// token.count == -1 TODO
// token.count == -2 TODO
type token struct {
	symbol rune
	count  int
}

func (t token) String() string {
	count := t.count
	if t.count == -1 {
		count = 1 // if number of times to repeat rune hadn't been given
	}
	return strings.Repeat(string(t.symbol), count)
}

func DecodeString(s string) (string, error) {
	result := ""

	// if token finished, currentToken.count sets to -2
	var currentToken token = token{count: -2}
	esc := false

	for _, r := range s {
		switch {
		case r == '\\':
			// definitely start of new token,
			if currentToken.count != -2 {
				result += currentToken.String()
				currentToken.count = -2
			}

			// two consecutive "\" simbols, means `\`
			if esc {
				esc = false
				currentToken = token{
					symbol: r,
					count:  -1,
				}
			} else {
				esc = true
			}

		case r >= '0' && r <= '9':
			// number
			// if before was `\`, should treat digit as token symbol
			if esc {
				esc = false

				if currentToken.count != -2 {
					result += currentToken.String()
				}

				currentToken = token{
					symbol: r,
					count:  -1,
				}
				continue
			}

			// happens if digit is in first symbol of string
			if currentToken.count == -2 {
				return "", ErrIncorrectString
			}

			if currentToken.count == -1 {
				currentToken.count = 0 // tmp, so line 82 calculated correctly
			}

			newCount, err := strconv.Atoi(string(r))
			if err != nil {
				return "", errors.New("internal err: " + err.Error())
			}

			// if several digits consecutevly
			currentToken.count = currentToken.count*10 + newCount
		default:
			// every other symbol
			// definitely start of new token,
			if esc {
				return "", ErrIncorrectString
			}

			if currentToken.count != -2 {
				result += currentToken.String()
			}

			currentToken = token{
				symbol: r,
				count:  -1,
			}
		}
	}

	// if last symbol was unescaped `\`
	if esc {
		return "", ErrIncorrectString
	}

	if currentToken.count != -2 {
		result += currentToken.String()
	}

	return result, nil
}

func main() {
	s, err := DecodeString(`а4bc4\5`)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println(s)
}
