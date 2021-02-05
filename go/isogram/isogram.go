package isogram

import (
	"strings"
)

//IsIsogram checks word for repeating runes
func IsIsogram(word string) bool {

	countMap := make(map[rune]bool)

	for _, c := range strings.ToLower(word) {
		if c == ' ' || c == '-' {
			continue
		}
		if countMap[c] {
			return false
		}
		countMap[c] = true
	}

	return true

}
