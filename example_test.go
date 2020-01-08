package must_test

import (
	"fmt"

	"github.com/gregoryv/must"
)

func ExampleCatchPanic() {
	err := myfunc()
	fmt.Println(err)
	// output: failed
}

func myfunc() (err error) {
	defer must.CatchPanic(&err)
	panic("failed")
}
