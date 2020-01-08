package must_test

import (
	"fmt"

	"github.com/gregoryv/must"
)

func ExampleCatchPanic() {
	err := func() (err error) {
		defer must.CatchPanic(&err)
		must.OpenFile("/no_such_file")
		return
	}()
	fmt.Println(err)
	// output: open /no_such_file: no such file or directory
}

func ExampleOpenFile() {
	fh := must.OpenFile("example_test.go")
	fh.Close()
}
