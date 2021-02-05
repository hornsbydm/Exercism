package hamming

import "errors"

//Distance calculates the hamming distance between 2 DNA sequences.
func Distance(a, b string) (retVal int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands not of same length")
	}

	for i, _ := range a {
		if a[i] != b[i] {
			retVal++
		}
	}
	return retVal, err
}
