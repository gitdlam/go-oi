package oi


import (
	"io"
)


func WriteNopCloser(w io.Writer) io.WriteCloser {
	wc := internalWriteNopCloser{
		writer:w,
	}

	return &wc
}


type internalWriteNopCloser struct {
	writer io.Writer
}


func (wc * internalWriteNopCloser) Write(p []byte) (n int, err error) {
	return wc.writer.Write(p)
}


func (wc * internalWriteNopCloser) Close() error {
	return nil
}
