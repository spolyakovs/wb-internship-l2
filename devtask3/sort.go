package main

import "bytes"

// sortLines sorts slice of line INPLACE with quicksort
func sortLines(lines []string) {
	values := valuesFromLines(lines)
	quicksortLines(lines, values)
}

func quicksortLines(lines []string, values [][]byte) {
	// values are []byte to compare, considering kFlag and nFlag
	if len(lines) < 2 {
		return
	}

	pivotIndex := len(lines) - 1 // select last element as pivot to divide array
	leftMostIndex := 0

	for i := 0; i < len(lines); i++ {

		if lessValues(values[i], values[pivotIndex]) {
			// move everything less, than pivot to left side
			lines[leftMostIndex], lines[i] = lines[i], lines[leftMostIndex]
			values[leftMostIndex], values[i] = values[i], values[leftMostIndex]
			leftMostIndex++
		}
	}

	// after cycle leftMostIndex will point to first elem, greater than pivot
	lines[leftMostIndex], lines[pivotIndex] = lines[pivotIndex], lines[leftMostIndex]
	values[leftMostIndex], values[pivotIndex] = values[pivotIndex], values[leftMostIndex]

	quicksortLines(lines[:leftMostIndex], values[:leftMostIndex])
	quicksortLines(lines[leftMostIndex+1:], values[leftMostIndex+1:])
}

// lessValues returns if value1 < value2 (> if rFlag == true)
func lessValues(value1 []byte, value2 []byte) bool {
	compare := bytes.Compare(value1, value2)

	if rFlag {
		compare = -compare // reverse comparing "direction"
	}

	return compare == -1
}
