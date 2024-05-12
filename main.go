package main

/**
 * main function is the entry point of the program.
 * It calls the recursivePow function with arguments 3 and 2,
 * and then prints the result.
 */
func main() {
	/**
	 * Println function prints the result of recursivePow function.
	 * It takes no arguments and returns nothing.
	 */
	println(recursivePow(3, 2))
}

/**
 * recursivePow function calculates the power of a number using recursion.
 * It takes two arguments: x (the base of the power) and n (the exponent).
 * The function returns the result of the power calculation.
 */
func recursivePow(x int, n int) int {
	/**
	 * If n is 0, the function returns 1 as the power of any number to the power of 0 is 1.
	 */
	if n == 0 {
		return 1
	}

	/**
	 * If n is not 0, the function recursively calls itself with the arguments x and n-1,
	 * and then multiplies the result by x.
	 */
	return x * recursivePow(x, n-1)
}
