package diffsquares

//SquareOfSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares returns thw sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference returns the difference between the square of the sum and the sum of squares of the first n natural numbers.
func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)
}
