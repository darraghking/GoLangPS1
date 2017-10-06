package main

//  "time" gets the current time, date and location
import (
	"fmt"
	"time"
)

func main() {

	//Initiate time
	t := time.Now()
	
	// The '.Hour' and '.Minute' after 't' allows you to get whatever part of the time import you need 
	fmt.Println("The time is \n", t.Hour(), ":", t.Minute())

	// Same as the above line, it works for the date as well as the time
	fmt.Println("The date is \n", t.Day(), "/", t.Month(), "/", t.Year())

}