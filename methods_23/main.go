package main

import (
	//	"fmt"
	"io"
	//	"io/ioutil"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const (
	rotateBy    = 13  // rotateBy refers to rotation index
	lowerMin    = 97  // lowerMin ASCII('a')
	lowerMax    = 122 // lowerMax ASCII('z')
	upperMin    = 65  // upperMin ASCII('A')
	upperMax    = 90  // upperMax ASCII('Z')
	letterCount = 26  // letterCount refers to total alphabet count
)

//rot13 - Rotates by playing with ASCII equivalent of chars
func rot13(b byte) byte {
	c := b
	if b >= lowerMin && b <= lowerMax {
		c = ((b-lowerMin)+rotateBy)%letterCount + lowerMin
	} else if b >= upperMin && b <= upperMax {
		c = ((b-upperMin)+rotateBy)%letterCount + upperMin
	}
	return c
}

func (R rot13Reader) Read(p []byte) (n int, err error) {
	n, err = R.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

/*
func (R rot13Reader)String() string {
	var data string
	if b, err := ioutil.ReadAll(R.r); err == nil {
    	data = string(b)
	}
	return data
}
*/

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	//fmt.Println("Input:", r)
	//fmt.Println("Input:", r) //This is printing empty. ioutils.ReadAll empties the io.Reader??
	io.Copy(os.Stdout, &r)

}
