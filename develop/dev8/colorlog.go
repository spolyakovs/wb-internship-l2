package main

import "strings"

func redBoldString(s string) string {

	return addColor(s, redBoldColor)
}

func greenBoldString(s string) string {

	return addColor(s, greenBoldColor)
}

func blueBoldString(s string) string {

	return addColor(s, blueBoldColor)
}

func addColor(s string, color string) string {
	sb := strings.Builder{}

	sb.WriteString(color)
	sb.WriteString(s)
	sb.WriteString(resetColor)

	return sb.String()
}
