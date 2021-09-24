package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r * rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		c := b[i]
		if 'A' <= c && c <= 'Z' {
			b[i] = (c - 'A' + 13) % 26 + 'A'
		} else if 'a' <= c && c <= 'z' {
			b[i] = (c - 'a' + 13) % 26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
