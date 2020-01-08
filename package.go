package must

import (
	"os"
)

// CatchPanic sets err if a panic is recovered
func CatchPanic(err *error) {
	e := recover()
	switch e := e.(type) {
	case nil:
	case error:
		*err = e
	}
}

// OpenFile opens filename or panics
func OpenFile(filename string) *os.File {
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return fh
}
