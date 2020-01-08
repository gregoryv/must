package must

import "fmt"

// CatchPanic sets err if a panic is recovered
func CatchPanic(err *error) {
	e := recover()
	switch e := e.(type) {
	case nil:
	case error:
		err = &e
	default:
		*err = fmt.Errorf("%v", e)
	}
}
