package ascii

func ContainNonPrintChar(words string) bool {
	for _, char := range words {
		if char < 32 || char > 126 {
			return true
		}
	}
	return false
}
