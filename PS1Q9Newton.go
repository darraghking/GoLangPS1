package main

import (
	"fmt"
	"math"
)

func z_next(z float64, x float64) float64 {
	return z - (((z * z) - x) / (2 * z))
}

func main() { 
	// The number to find the square root of
	x := 64.0

	// The initial guess
	z := float64(1)

	// Iterate until the next guess is the same as the current guess
	for z = 1.0 ; z!= z_next(z,x) ; z = z_next(z,x) {
		// Print the guess on each iteration
		fmt.Println("\nCurrent guess: %f ", z)
	}	

	// Finally, z is a good approximation of the square root of x 
	//fmt.Printf("The square root of %.8f is %f.\n",x,z)
	fmt.Printf("The square root of ", x," is ",z,".\n")
	// Print out the math.Sqrt value
	fmt.Printf("math.Sqrt gives the value as %f.\n", math.Sqrt(x))

}