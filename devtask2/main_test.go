package main

import (
	"errors"
	"testing"
)

func TestDecodeString(t *testing.T) {
	t.Run("Empty string", templateTest(t, "", ""))

	t.Run("Without repeat", templateTest(t, "abcd", "abcd"))

	t.Run("Without repeat digit", templateTest(t, "abcd", "abcd"))

	t.Run("With unicodes", templateTest(t, "а1б2в3", "аббввв"))

	t.Run("With 0 repeat", templateTest(t, "a0b0c", "c"))

	t.Run("With multuple digits", templateTest(t, `\0ab02c12d`, "0abbccccccccccccd"))

	t.Run("With escape character", templateTest(t, `ab\0c\\`, `ab0c\`))

	t.Run("With miltiple escape characters", templateTest(t, `ab\\2`, `ab\\`))

	t.Run("With last escape character", templateTestErr(t, `abcd\`))

	t.Run("Incorrect escape", templateTestErr(t, `\abcd`))

	t.Run("First digit", templateTestErr(t, `3abcd`))
}

func templateTest(t *testing.T, rawString string, wantString string) func(t *testing.T) {
	return func(t *testing.T) {
		gotString, err := DecodeString(rawString)
		if err != nil {
			t.Errorf("error: %v", err)
			return
		}
		if gotString != wantString {
			t.Errorf("decoded incorrectly\n\tgot: %v\n\twanted: %v", gotString, wantString)
		}
	}
}

func templateTestErr(t *testing.T, rawString string) func(t *testing.T) {
	return func(t *testing.T) {
		gotString, err := DecodeString(rawString)
		if err == nil {
			t.Errorf("decoded incorrect string\n\tgot: %v", gotString)
			return
		}
		if !errors.Is(err, ErrIncorrectString) {
			t.Errorf("got wrong error\n\tgot: %v\n\twanted: incorrect string", err)
		}
	}
}
