package main

import "golang.org/x/tour/pic"

// Pic - function to return the data for pic.Show function
func Pic(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)

	for i := range data {
		data[i] = make([]uint8, dx)
		for j := range data[i] {
			data[i][j] = uint8((i+j)/2 + i ^ j - i/(j+1))
		}
	}

	return data
}

func main() {
	pic.Show(Pic)
}
