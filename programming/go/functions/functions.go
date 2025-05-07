package main

import "fmt"

func plus(m int, n int) int {
	return m + n
}
func plusPlus(m, n, p int) int {
	return m + n + p
}
func main() {

	res := plus(7, 12)
	fmt.Println("7 + 12 =", res)

	res = plusPlus(7, 12, 5)
	fmt.Println("7 + 12 + 5 =", res)

}
