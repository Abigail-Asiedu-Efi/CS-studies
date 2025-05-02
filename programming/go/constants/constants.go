package main

import (
	"fmt"
	"math"
)

const y string = "constant"

func main() {
	fmt.Println(y)

	const n = 500000000
	const x = 3e20 / n

	fmt.Println(x)
	fmt.Println(int64(x))
	fmt.Println(math.Sin(n))
}
