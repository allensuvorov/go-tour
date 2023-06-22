package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", fmt.Sprint(float64(e)))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	keepGoing := true
	prev := 0.0
	for i := 0; keepGoing && i < 100 ; i++ { 
		z -= (z*z - x) / (2*z)
		if z == prev {
			keepGoing = false
			fmt.Println(i)
		} 
		if math.Abs(z*z - x) < x/math.Pow(10, 40) {
			keepGoing = false
			fmt.Println(i, z*z)
		}

		prev = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
