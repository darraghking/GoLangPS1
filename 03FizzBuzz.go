package main

import "fmt"

func main() {
	
	var count int = 0

	for count<100{
		count++
	   if count % 3 == 0 && count % 5 ==0 {
		   fmt.Println("FizzBuzz")
       } else if count % 3 == 0 {
		   fmt.Println("Fizz")
	   } else if count % 5 == 0 {
		   fmt.Println("Buzz")
	   } else {
		   fmt.Println(count)
	   }// End if else

	}// End for count<=100

}// End main
