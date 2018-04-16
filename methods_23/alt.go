package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var lowerCase = []byte("abcdefghijklmnopqrstuvwxyz") // lowerCase byte array of lower case letters
var upperCase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ") // upperCase byte array of upper case letters

const (
	rotateBy    = 13 // rotateBy refers to rotation index
	letterCount = 26 // letterCount refers to total alphabet count
)

//rot13 - Rotates by playing with the index of initial byte array
func rot13(b byte) byte {
	i := bytes.IndexByte(lowerCase, b)
	if i != -1 {
		return lowerCase[(i+rotateBy)%letterCount]
	}

	i = bytes.IndexByte(upperCase, b)
	if i != -1 {
		return upperCase[(i+rotateBy)%letterCount]
	}

	return b
}

func (R rot13Reader) Read(p []byte) (n int, err error) {
	n, err = R.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
