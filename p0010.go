package main

/*
 * Summation of primes
 * Problem 10
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 *
 * Find the sum of all the primes below two million.
 */

import "fmt"
//import "math"
import "math/big"

func main() {
	var sum int

	for i := 2; i < 2000000; i++ {
		var bi big.Int
		bi.SetUint64( uint64( i ) )
		if( bi.ProbablyPrime( 20 ) ) {
			sum += i
		}
	}
	
	fmt.Printf( "sum of primes below 2000000: %d\n", sum )
}

