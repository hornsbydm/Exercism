package grains

import "errors"

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
	return 0xFFFF_FFFF_FFFF_FFFF
}

// Square returns the number of grains on a given square (1-based).
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("input value not in domain")
	}
	return 1 << (n - 1), nil
}
