package cipher

import (
	"io"
)

/*
Task: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') { // Tests between the 13 first places in alphabet
			p[i] = p[i] + 13 // Substitues 13 places in alphabet.
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') { // Tests between the 13 last places in alphabet
			p[i] = p[i] - 13 // Substitues 13 places in alphabet.
		}
	}
	return n, err
}
func test() {}
