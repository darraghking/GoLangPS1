package main

import (
   "fmt"
   "math/big"
)

// Function to calculate the factorial
func factorial(n int64) *big.Int {
   if n < 0 {
      return big.NewInt(1)
   }
   if (n==0) {
      return big.NewInt(1)
   }
   bigN := big.NewInt(n)
   return bigN.Mul(bigN, factorial(n-1))
}

// Function that gets the sum of the factorial for 100!
func add(number *big.Int) *big.Int {
    ten := big.NewInt(10)
    sum := big.NewInt(0)
    mod := big.NewInt(0)
    for ten.Cmp(number)<0 {
      sum.Add(sum, mod.Mod(number,ten))
      number.Div(number,ten)
    }
    sum.Add(sum,number)
  return sum
}

func main() {
	// Prints the results
	fmt.Print("The Factorial of 100 is: ")
	fmt.Print(factorial(100))
	fmt.Print("\nThe sum of the Factorial of 100 is: ")
	fmt.Print(add(factorial(100)))
	
}