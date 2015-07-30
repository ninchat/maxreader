package maxreader

import (
	"errors"
	"io"
)

// ErrReadLimit is the interesting error.
var ErrReadLimit = errors.New("read limit exceeded")

type maxReader io.LimitedReader

// New creates a reader which behaves mostly like io.LimitedReader.  Its Read
// method returns an error if the underlying reader produces surplus data.
func New(r io.Reader, max int64) io.Reader {
	return maxReader(io.LimitedReader{
		R: r,
		N: max + 1,
	})
}

func (mr maxReader) Read(buf []byte) (n int, err error) {
	lr := io.LimitedReader(mr)
	n, err = lr.Read(buf)
	if lr.N <= 0 {
		err = ErrReadLimit
	}
	return
}
