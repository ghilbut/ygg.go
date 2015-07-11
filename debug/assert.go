package assert

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

const kFormat = `
>>> [ygg.go] assert.%s >>>
[FILE] %s:%d\n
[FUNC] %s
%s
<<< [ygg.go] assert.%s <<<`

func assert(t string, a ...interface{}) {
	// NOTE(ghilbut):
	// http://stackoverflow.com/questions/25927660/golang-get-current-scope-of-function-name

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	message := fmt.Sprint(a...)
	log.Panicf(kFormat, t, file, line, f.Name(), message, t)
}

func True(condition bool, a ...interface{}) {
	if !condition {
		assert("True", a)
	}
}

func False(condition bool, a ...interface{}) {
	if condition {
		assert("False", a)
	}
}

func Nil(condition interface{}, a ...interface{}) {
	if condition != nil && !reflect.ValueOf(condition).IsNil() {
		assert("Nil", a)
	}
}

func NotNil(condition interface{}, a ...interface{}) {
	if condition == nil || reflect.ValueOf(condition).IsNil() {
		assert("NotNil", a)
	}
}

func Contains(m interface{}, k interface{}, a ...interface{}) {

	const kName = "Contains"

	if m == nil {
		assert(kName, "m should not be nil.", "\n", a)
	}
	if k == nil {
		assert(kName, "k should not be nil.", "\n", a)
	}

	vm := reflect.ValueOf(m)
	vk := reflect.ValueOf(k)

	if vm.Kind() != reflect.Map {
		assert(kName, "m should be map.", "\n", a)
	}
	if vm.Type().Key() != vk.Type() {
		assert(kName, "k is not allowed key type.", "\n", a)
	}

	if vm.MapIndex(vk).Interface() == nil {
		assert(kName, a)
	}
}
