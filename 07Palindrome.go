package main

import (
 "fmt"
 "strings"
)

func main() {

 var ip string
 fmt.Println("Enter string:")
 fmt.Scanf("%s\n", &ip)
 ip = strings.ToLower(ip)
 fmt.Println(isP(ip))
}

// Function to test if the string entered is a Palindrome
func isP(s string) string {

// Mid takes the length of the string and halfs it to get the middle,
// Last finds the last letter of the string
 mid := len(s) / 2
 last := len(s) - 1

 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "This is not a Palindrome"
  }
 }
 return "This IS a Palindrome"
}