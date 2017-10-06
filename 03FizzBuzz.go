package main

import "fmt"

func main() {
	
	var count int = 0

	for count<100{

		// Count++ gets each number in the loop
		count++

		// First line of the if prints "FizzBuzz" each time a number in the loop is divisible by 3 and 5.
		// Then the next 2 lines of the if statement print "Fizz" for a number that is divisble by 3,
		// and "Buzz" for a number that is divisible 5. 
		// 'fmt.Println(count)' prints every number, including the numbers that have been changed
	   if count % 3 == 0 && count % 5 ==0 {
		   fmt.Println("FizzBuzz")
       } else if count % 3 == 0 {
		   fmt.Println("Fizz")
	   } else if count % 5 == 0 {
		   fmt.Println("Buzz")
	   } else {
		   fmt.Println(count)
	   }// End if else

	}// End for count<100

}// End main
