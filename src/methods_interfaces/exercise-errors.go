package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	fmt.Sprint(float64(e)) 
	return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	result, err := Sqrt(2)
	fmt.Println(result, err)

	result, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
