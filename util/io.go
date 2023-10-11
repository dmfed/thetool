package util

import (
	"io"
	"os"
)

func OpenInput(src string) (io.ReadCloser, error) {
	if src == "" {
		return io.NopCloser(os.Stdin), nil
	}
	f, err := os.Open(src)
	return f, err
}

func OpenOutput(dst string) (io.WriteCloser, error) {
	if dst == "" {
		return os.Stdout, nil
	}
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	return f, err
}
