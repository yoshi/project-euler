package main

/* 
 * Largest palindrome product
 * Problem 4
 
 * A palindromic number reads the same both ways. The largest
 * palindrome made from the product of two 2-digit numbers is 9009 =
 * 91 × 99.

 * Find the largest palindrome made from the product of two 3-digit numbers.
 */

import "fmt"
import "strconv"
//import "sort"
//import "strings"

func reverse( s string ) string {
	a := []rune( s )
	b := make( []rune, len(a) )

	j := len(a) - 1

	for i,_ := range( a ) {
		b[j] = a[i]
		j -= 1
	}

	t := string(b)

	return t;
}

func isPalindrome( p int ) bool {
	rp, _ := strconv.ParseInt( reverse( strconv.Itoa( p ) ), 0, 0 )
	if p == int(rp) {
		return true
	}
	return false
}

func main() {
	base := 999
	candidate, m, n := 0, 0, 0

	for i := base; i > 0; i-- {
		for j := base; j > 0; j-- {
			p := i * j
			if isPalindrome( p ) {
				if p > candidate {
					candidate = p
					m,n = i,j
				}
			}
		}
	}

	fmt.Printf( "For a base %d, %d (%d * %d) is the largest palindrome number\n", base, candidate, m, n )
}
