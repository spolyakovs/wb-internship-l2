package main

type Line struct {
	str string
	num int
}

func NewLine(str string, num int) *Line {
	return &Line{
		str: str,
		num: num,
	}
}
