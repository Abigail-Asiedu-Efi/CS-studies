package main

import "fmt"

func main() {
	if 9%2 == 0 {
		fmt.Println("9 is an even number")
	} else {
		fmt.Println("9 is an odd number")
	}
	if 6%3 == 0 {
		fmt.Println("6 is divisible by 3")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 is even")
	}
	if num := 9; num < 0 {
		fmt.Println("Num is a negative value")
	} else if num < 10 {
		fmt.Println("Num is a single digit")
	} else {
		fmt.Println("Num is a multiple digit")
	}
}
