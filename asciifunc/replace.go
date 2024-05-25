package ascii

import "strings"

func ReplaceNonPrintChar(words string) string {
	words = strings.ReplaceAll(words, "\\t", "    ")
	words = strings.ReplaceAll(words, "\\a", "\a")
	words = strings.ReplaceAll(words, "\\b", "\b")
	words = strings.ReplaceAll(words, "\\f", "\f")
	words = strings.ReplaceAll(words, "\\r", "\r")
	words = strings.ReplaceAll(words, "\\v", "\v")
	words = strings.ReplaceAll(words, "\n", "\\n")
	
	
	return words
}
