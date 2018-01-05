package main

import "fmt"
import (
	"./chapter1"
	"strings"
)

func main() {
	exercise1_1()
	exercise1_2()
}


func exercise1_1(){
	fmt.Printf("Exercise 1.1 - hasAllUniqueCharacters(\"abcde\"): %t\n", chapter1.HasAllUniqueCharacters("abcde")) //true
	fmt.Printf("Exercise 1.1 - hasAllUniqueCharacters(\"abcdbe\"): %t\n", chapter1.HasAllUniqueCharacters("abcdbe")) //false
	fmt.Printf("Exercise 1.1 - hasAllUniqueCharacters(\"abecde\"): %t\n", chapter1.HasAllUniqueCharacters("abecde")) //false
	fmt.Printf("Exercise 1.1 - hasAllUniqueCharacters(\"aabcde\"): %t\n", chapter1.HasAllUniqueCharacters("aabcde")) //false

	fmt.Printf("Exercise 1.1+ - hasAllUniqueCharacters2(\"abcde\"): %t\n", chapter1.HasAllUniqueCharacters2(strings.Split("abcde", ""))) //true
	fmt.Printf("Exercise 1.1+ - hasAllUniqueCharacters2(\"abcbde\"): %t\n", chapter1.HasAllUniqueCharacters2(strings.Split("abcbde", ""))) //false
	fmt.Printf("Exercise 1.1+ - hasAllUniqueCharacters2(\"abecdee\"): %t\n", chapter1.HasAllUniqueCharacters2(strings.Split("abecdee", ""))) //false
	fmt.Printf("Exercise 1.1+ - hasAllUniqueCharacters2(\"axagbsbcde\"): %t\n", chapter1.HasAllUniqueCharacters2(strings.Split("axagbsbcde", ""))) //false
}

func exercise1_2()  {
	fmt.Print("Exercise 1.2 - reverseCString([]int32{'h','e', 'l','l','o'}): %v\n", chapter1.ReverseCString([]string{"h","e", "l","l","o"})) //olleh
}