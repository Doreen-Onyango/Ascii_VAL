package main

import (
	"fmt"
	"os"

	ascii "ascii/asciifunc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run. <input string>")
		return
	}

	words := os.Args[1]

	words = ascii.ReplaceNonPrintChar(words)
	if ascii.ContainNonPrintChar(words) {
		fmt.Println("contains non-printable characters")
		return
	}
	if ascii.SpecialCases(words) {
		return
	}
	ascii.ProcessWords(words)
}
