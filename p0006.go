package main

/* 
 * Sum square difference
 * Problem 6
 *
 * The sum of the squares of the first ten natural numbers is,
 *
 * 1^2 + 2^2 + ... + 10^2 = 385
 *
 * The square of the sum of the first ten natural numbers is,
 *
 * (1 + 2 + ... + 10)^2 = 55^2 = 3025
 *
 * Hence the difference between the sum of the squares of the first
 * ten natural numbers and the square of the sum is 3025 − 385 = 2640.
 *
 * Find the difference between the sum of the squares of the first one
 * hundred natural numbers and the square of the sum.
 */
 
import "fmt"

func main() {
	var susq, sqsu int
	var max int = 100

	for i := 1; i <= max; i++ {
		susq += i * i
		sqsu += i
	}

	sqsu *= sqsu

	fmt.Printf( "for a max %d, sqsu %d - susq %d = %d\n", max, sqsu, susq, sqsu - susq )
}
