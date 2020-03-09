package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err == io.EOF {
		return n, err
	}
	for i, v := range b[:n] {
		if v >= 65 && v <= 90 {
			b[i] = ((v+13)-65)%26 + 65
		} else if v >= 97 && v <= 122 {
			b[i] = ((v+13)-97)%26 + 97
		}
		continue
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
