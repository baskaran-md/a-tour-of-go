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

func manSqrt(x float64) (float64, int) {
	z := x / 2
	var result float64
	i := 1
	for ; i < 20; i++ {
		//fmt.Println("Loop:", i)
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("RoundOff: (%v)==(%v)\n",result, roundOff(result, precision))
		if roundOff(result, precision) == roundOff(z, precision) {
			//fmt.Println("Same result @ Loop:", i)
			return result, i
		}
		result = z
	}
	return result, i
}

func main() {
	x := 2.0
	fmt.Printf("Sqrt(%v)=%f (Math)\n", x, Sqrt(x))
	result, loop := manSqrt(x)
	fmt.Printf("Sqrt(%v)=%f (Custom)[Loops: %d]\n", x, result, loop)

	x = 20.0
	fmt.Printf("Sqrt(%v)=%f (Math)\n", x, Sqrt(x))
	result, loop = manSqrt(x)
	fmt.Printf("Sqrt(%v)=%f (Custom)[Loops: %d]\n", x, result, loop)
}
