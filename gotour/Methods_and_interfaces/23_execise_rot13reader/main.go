package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	rot13(b)
	return n, err
}

func rot13(b []byte) {
	for i, v := range b {
		// A - M
		// a - m
		if (v >= 'a' && v <= 'm') || (v >= 'A' && v <= 'M') {
			b[i] = v + 13
		}
		
		// M - Z
		// m - z
		if (v >= 'n' && v <= 'z') || (v >= 'N' && v <= 'Z') {
			b[i] = v - 13
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // *strings.Reader
	r := rot13Reader{s} // instance of rot13Reader is created
	io.Copy(os.Stdout, &r) // here Read method is fired
}
