package util

import (
	"bufio"
	"context"
	"io"
)

type Reader struct {
	scanner *bufio.Scanner
	err     error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{scanner: bufio.NewScanner(r)}
}

func (r *Reader) Read(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for r.scanner.Scan() {
			line := r.scanner.Text()
			if line == "" {
				continue
			}
			select {
			case <-ctx.Done():
				r.err = ctx.Err()
				return
			case ch <- r.scanner.Text():
			}

		}
		r.err = r.scanner.Err()
	}()
	return ch
}

func (r *Reader) Err() error {
	return r.err
}
