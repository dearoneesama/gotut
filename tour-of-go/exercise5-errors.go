package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt: %v", float64(e))
}

const epsilon = 1e-14

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for Abs(z*z - x) >= epsilon  {
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
