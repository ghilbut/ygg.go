package debug_test

import (
	. "github.com/ghilbut/ygg.go/debug"
	"testing"
)

func Test_assert_True_is_true(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	Assert(true, "this is assert test. this message is not allowed.")
}

func Test_assert_True_is_false(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	Assert(false, "test message. this message should be displayed.")
}
