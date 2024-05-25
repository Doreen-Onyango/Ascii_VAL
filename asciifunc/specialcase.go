package ascii

import "fmt"

func SpecialCases(words string) bool {
	if words == "\\n" {
		fmt.Println()
		return true
	} else if words == "" {
		return true
	}
	return false
}
