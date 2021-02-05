// Package scrabble provides functions for the scrabble game
package scrabble

// Score tallies the letter scores for a scrabble word
func Score(word string) (result int) {
	for _, letter := range word {
		if letter >= 'a' {
			letter -= ' '
		}
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}
	}
	return
}
