package must

import (
	"fmt"
	"testing"
)

func TestCatchPanic(t *testing.T) {
	ok := func(panicValue interface{}) {
		t.Helper()
		t.Run("", func(t *testing.T) {
			var err error
			defer func(err *error) {
				if err == nil {
					t.Errorf("panic(%T) should fail", panicValue)
				}
			}(&err)
			defer CatchPanic(&err)
			panic(panicValue)
		})
	}
	ok(1)
	ok(fmt.Errorf("any"))
}

func TestCatchPanic_no_error(t *testing.T) {
	var err error
	defer func() {
		if err != nil {
			t.Error("should fail")
		}
	}()
	defer CatchPanic(&err)
}

func Test_ok_funcs(t *testing.T) {
	OpenFile("unit_test.go").Close()
}
