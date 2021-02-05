// Package luhn provides implementation for validing luhn numbers.
package luhn

import (
	"errors"
	"strings"
)

// Valid checks whether provided number is valid luhn.
func Valid(in string) bool {

	parsedInput, err := makeSlice(in)

	if err != nil {
		return false
	}

	if len(parsedInput) <= 1 { //Strings of length 1 or less are not valid.
		return false
	}

	return doubleAndSum(parsedInput)%10 == 0
}

// doubleAndSum returns the sum, doubling every second digit from the right squared
func doubleAndSum(in []int) (retVal int) {

	for i, x := range in {
		if (len(in)-i)%2 != 0 { //Don't double just sum
			retVal += x
			continue
		}

		if x*2 > 9 { //Normalize to base 9
			retVal += x*2 - 9
			continue
		}
		retVal += x * 2 //double
	}

	return
}

// makeSlice makes integer slice from input string.
func makeSlice(in string) (retVal []int, err error) {
	for _, x := range strings.ReplaceAll(in, " ", "") {
		if x < '0' || x > '9' {
			return nil, errors.New("invalid characters in input string")
		}
		retVal = append(retVal, int(x-'0'))
	}
	return
}
