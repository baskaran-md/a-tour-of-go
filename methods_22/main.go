package main

import "golang.org/x/tour/reader"
import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})

	var r MyReader
	b := make([]byte, 10)
	len, err := r.Read(b)
	fmt.Println(string(b), len, err)
}
