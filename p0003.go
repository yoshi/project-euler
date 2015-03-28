package main

/*
 * Largest prime factor
 * Problem 3
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143 ?
 */

import "fmt"
import "math"
import "math/big"
import "sort"
import "github.com/cznic/sortutil"

func gp( slice sortutil.Uint64Slice ) uint64 {
	for i := slice.Len() / 2 - 1; i >= 0; i-- {
		var bi big.Int
		bi.SetUint64( slice[i] )
		if bi.ProbablyPrime( 20 ) {
			return slice[i]
		}
	}

	return 0
}

func factors( a uint64 ) []uint64 {
	var i uint64
	
	slice := make( sortutil.Uint64Slice, 0 )

	for i = 1; i <= uint64( math.Sqrt( float64( a ) ) ); i++ {
		if a % i == 0 {
//			fmt.Printf( "%d and %d are factors of %d\n", i, a/i , a )
			slice = append( slice, i )
			slice = append( slice, a/i )
		}
	}

	sort.Sort( slice )

	return slice
}

func main() {
	vs := [2]uint64 { 13195, 600851475143 }
	for _, v := range vs {
		fmt.Printf( "The greatest prime factor of %d is %d\n", v, gp( factors( v ) ) )		
	}
}
