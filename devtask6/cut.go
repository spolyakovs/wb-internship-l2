package main

import (
	"fmt"
	"strings"
)

func Cut(line string, del string) {
	if del == "" || !strings.Contains(line, del) {
		if !sFlag {
			fmt.Println(line)
		}
		return
	}

	fields := strings.Split(line, del)
	// fmt.Println("Fields len:", len(fields))

	if fFlag > len(fields) {
		fmt.Println()
	} else {
		fmt.Println(fields[fFlag-1])
	}
}
