package main

import (
	"fmt"
)

const epsilon = 1e-14

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) (z float64) {
	z = 1.0
	for Abs(z*z - x) >= epsilon  {
		z -= (z*z - x) / (2*z)
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
}
