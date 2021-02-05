package reverse

//Reverse reverses a string
func Reverse(s string) string {
	chars := make([]rune, len(s))

	for i, c := range s {
		chars[len(s)-i-1] = c
	}

	return string(chars)

}
