package main

import (
	"fmt"
	"math"
)

// Sqrt - Performs math.Sqrt function
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func roundOff(x float64, d int) float64 {
	return math.Round(x*(math.Pow(10, float64(d)))) / math.Pow(10, float64(d))
}

const precision = 5

// ErrNegativeSqrt - Error handling for negative numbers
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func manSqrt(x float64) (float64, int, error) {
	if x < 0.0 {
		return 0, 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	var result float64
	i := 1
	for ; i < 20; i++ {
		//fmt.Println("Loop:", i)
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("RoundOff: (%v)==(%v)\n",result, roundOff(result, precision))
		if roundOff(result, precision) == roundOff(z, precision) {
			//fmt.Println("Same result @ Loop:", i)
			return result, i, nil
		}
		result = z
	}
	return result, i, nil
}

func calcSqrt(x float64) {
	fmt.Printf("Sqrt(%v)=%f (Math)\n", x, Sqrt(x))
	result, loop, err := manSqrt(x)
	if err != nil {
		fmt.Printf("Error while calculating sqrt for %v. Message: %v\n", x, err)
	} else {
		fmt.Printf("Sqrt(%v)=%f (Custom)[Loops: %d]\n", x, result, loop)
	}
}

func main() {
	calcSqrt(2.0)
	calcSqrt(20.0)
	calcSqrt(-20.0)
	calcSqrt(49)
}
