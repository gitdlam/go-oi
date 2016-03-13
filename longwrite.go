package oi


import (
	"io"
)


const (
	lenBuffer = 64
)


// LongWrite trys to write the bytes from 'p' to the writer 'w', such that it deals
// with "short writes" where w.Write would return an error of io.ErrShortWrite and
// n < len(p).
func LongWrite(w io.Writer, p []byte) (int64, error) {

	numWritten := int64(0)
	for {
//@TODO: Should check to make sure this doesn't get stuck in an infinite loop writting nothing!
		n, err := w.Write(p)
		if nil != err && io.ErrShortWrite != err {
			return int64(n), err
		}
		numWritten += int64(n)

		if !(n < len(p)) {
			break
		}

		p = p[n:]

		if len(p) < 1 {
			break
		}
	}

	return numWritten, nil
}
