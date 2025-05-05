package main

import "fmt"

func main() {
	var m [5]int
	fmt.Println("emp: ", m)

	m[4] = 100
	fmt.Println("set: ", m)
	fmt.Println("get: ", m[4])

	fmt.Println("length: ", len(m))

	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", n)

	n = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", n)

	n = [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", n)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
