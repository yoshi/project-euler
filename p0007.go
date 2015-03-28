package main

/* 
 * 10001st prime
 * Problem 7
 *
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we
 * can see that the 6th prime is 13.
 *
 * What is the 10 001st prime number?
 */
 
import "fmt"
//import "math"
import "math/big"

func main() {
	count := 10001
	found := 0

	p := 2

	for found < count {
		var bi big.Int
		bi.SetUint64( uint64( p ) )
		if bi.ProbablyPrime( 50 ) {
			found++
			fmt.Printf( "%d th prime: %d\n", found, p )
		}
		p++
		if p % 2 == 0 {
			p++
		}
	}
}
