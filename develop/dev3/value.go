package main

import (
	"encoding/binary"
	"strconv"
	"strings"
	"unicode"
)

func valuesFromLines(lines []string) [][]byte {
	result := make([][]byte, len(lines))
	for i := range lines {
		result[i] = getValue(lines[i])
	}
	return result
}

// getValue returns []byte value for given line, considering kFlag and nFlag
func getValue(line string) []byte {
	result := make([]byte, 0)

	if kFlag != 0 {
		if column, ok := getColumn(line, kFlag-1); ok {
			result = append(result, getValueString(column)...)
		}
	}

	return append(result, getValueString(line)...)
}

// getValueNumber returns []byte value from string for given line
func getValueString(line string) []byte {
	if nFlag {
		value, ok := getValueNumber(line)
		if ok {
			return value
		}
	}
	// TODO: mCapFlag value
	// TODO: if bFlag trim tail whitespaces
	return []byte(line)
}

// getValueNumber returns []byte value from int for given line, if it can be converted to int
func getValueNumber(line string) (value []byte, ok bool) {
	// TODO: if hFlag parse "K", "G", "m" etc into 10^3, 10^6, 10^(-3) ...
	value = make([]byte, 8) // for uint64
	var num uint64

	num, err := strconv.ParseUint(line, 10, 0)
	if err != nil {
		return nil, false
	}

	binary.BigEndian.PutUint64(value, num)
	return value, true
}

// getColumn returns specified column from line if it exists
func getColumn(line string, columnNum int) (_ string, ok bool) {
	columns := strings.Split(stringsTrim(line), " ")
	if columnNum >= 0 && columnNum < len(columns) {
		return columns[columnNum], true
	}
	return "", false
}

// stringsTrim trims given string from multiple spaces(tabs, new lines etc) and replaces them with 1 " "
func stringsTrim(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	whitespaced := false
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			whitespaced = false
		} else {
			if !whitespaced {
				b.WriteRune(' ')
				whitespaced = true
			}
		}
	}
	return b.String()
}
