//This one was stolen
// https://stackoverflow.com/questions/31669266/tour-of-go-exercise-rot13reader

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rd rot13Reader) Read(b []byte) (int, error) {	
	m := make(map[byte]byte)
	input := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    output := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	for i := range input {
        m[input[i]] = output[i]
    }
	n, err := rd.r.Read(b)
	for i:= 0; i < n; i++ {
		if val, ok := m[b[i]]; ok {
			b[i] = val
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

}
