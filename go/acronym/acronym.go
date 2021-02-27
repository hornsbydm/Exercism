// Package acronym implements simple title shortining funcs
package acronym

import "strings"

// Abbreviate creates an acronym from a phrase.
func Abbreviate(s string) (retVal string) {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	for _, v := range strings.Fields(s) {
		retVal += string(v[0])
	}
	return strings.ToUpper(retVal)
}
