package main

/* 
 * Smallest multiple
 * Problem 5

 * 2520 is the smallest number that can be divided by each of the
 * numbers from 1 to 10 without any remainder.

 * What is the smallest positive number that is evenly divisible by
 * all of the numbers from 1 to 20?
 */
 
import "fmt"

func divTest( v int, count int ) bool {
	for i := 1; i <= count; i++ {
		if v % i != 0 {
			return false
		}
	}
	return true
}

func factorial( v int ) int {
	if( v == 0 ) {
		return 1
	}

	return v * factorial( v - 1 )
}

func main() {
	var count int = 20
	var upperbound int = factorial( count )

	for i := count; i <= upperbound; i++ {
		if divTest(i, count) {
			fmt.Printf( "count: %d upperbound: %d value: %d\n", count, upperbound, i )
			break
		}
	}
}
