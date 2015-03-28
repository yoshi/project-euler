package main

/*
 * Special Pythagorean triplet
 * Problem 9
 *
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 *
 * a^2 + b^2 = c^2
 *
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 */

import "fmt"
import "math"

// ok, I got a little carried away and made this super complicated by
// using trignometric functions.

// solves for the angle between a and b
func sss( a float64, b float64, c float64 ) float64 {
	gamma := math.Acos( ( a * a + b * b - c * c ) / ( 2 * a * b ) ) / (math.Pi / 180)

	return gamma
}

func aas( aAng float64, bAng float64, c float64 ) float64 {
	return 1.0
}

func angSolver() {
	// lets presume that a = 3, b = 4, and c = 5

	aAng := sss( 4, 5, 3 )
	bAng := sss( 5, 3, 4 )
	cAng := sss( 3, 4, 5 )

	fmt.Printf( "%f %f %f", aAng, bAng, cAng )
}

func losrt( c float64 ) (float64, float64) {
	
	x := math.Pi / 180

	fmt.Printf( "sin 30 %f\n", math.Sin(30 * x))
	fmt.Printf( "sin 60 %f\n", math.Sin(60 * x))
	fmt.Printf( "sin 90 %f\n", math.Sin(90 * x))

	a := ( c * math.Sin( 30 * x ) ) / math.Sin( 90 * x )
	b := ( c * math.Sin( 60 * x ) ) / math.Sin( 90 * x )

	return a, b
}

/*
so, geometricaly we know all the angles and we can use one side to make a projection of both of the other sides.  ASS is the only one that does not work, so AAS will be the best bet.

lets use the law of sines to determine the sides of the right triangle. 
*/
func oldMain() {
	return

	fmt.Printf( "huh: %f\n", math.Acos( 0 ) / (math.Pi / 180) )
	angSolver()

	// generate all squares for numbers below 500 as a guess
	sq := make( []int, 500)

	for i,_ := range sq {
		sq[i] = i*i
	}

	fmt.Printf( "len %d\n", len(sq) )

	c := 5.0
	a,b := losrt( c )
	fmt.Printf( "a: %f b: %f c: %f\n", a, b, c )
}

func main() {
	// In the end, I used brute force.

	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			for k := 0; k < 500; k++ {
				if( ( i*i + j*j == k*k ) && i+j+k == 1000) {
					fmt.Printf( "i: %d j: %d k: %d sum: %d product: %d\n", i, j, k, i + j + k, i*j*k )
					break
				}
			}
		}
	}
}

