package assert_test

import (
	. "github.com/ghilbut/ygg.go/debug"
	"testing"
)

const kDisplay = "[Test] this message should be displayed."
const kNotDisplay = "[Test] this message should not be displayed."

func Test_assert_True_is_false(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	True(false, kDisplay)
}

func Test_assert_True_is_true(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	True(true, kNotDisplay)
}

func Test_assert_False_is_true(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	False(true, kDisplay)
}

func Test_assert_False_is_false(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	False(false, kNotDisplay)
}

func Test_assert_Contains_when_check_invalid_type_key(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	target := make(map[string]bool)
	Contains(target, 1, kDisplay)
}

func Test_assert_Contains_when_not_contained(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	target := make(map[string]bool)
	Contains(target, "key", kDisplay)
}

func Test_assert_Contains_when_contained(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	target := make(map[string]bool)
	target["key"] = true
	Contains(target, "key", kNotDisplay)
}
