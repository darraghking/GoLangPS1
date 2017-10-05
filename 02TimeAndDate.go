package main

import (
	"fmt"
	"time"
)

func main() {
	
	t := time.Now()
	
	fmt.Println("The time is \n", t.Hour(), ":", t.Minute())

	fmt.Println("The date is \n", t.Day(), "/", t.Month(), "/", t.Year())

}