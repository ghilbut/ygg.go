package debug

import (
	"fmt"
	"log"
	"runtime"
)

const kFormat = `
>>>>>>>> ASSERT >>>>>>>>
[FILE] %s:%d\n
[FUNC] %s
%s
<<<<<<<< ASSERT <<<<<<<<`

func getDetail() (file string, line int, function string) {
	// NOTE(ghilbut):
	// http://stackoverflow.com/questions/25927660/golang-get-current-scope-of-function-name

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line = f.FileLine(pc[0])
	function = f.Name()
	return
}

func Assert(condition bool, a ...interface{}) {
	if !condition {
		file, line, function := getDetail()
		msg := fmt.Sprint(a...)
		log.Panicf(kFormat, file, line, function, msg)
	}
}

/*
func Assertf(condition bool, format string, a ...interface{}) {
	if !condition {
		file, line, function := getDetail()
		msg := fmt.Sprintf(format, a...)
		log.Panicf(kFormat, file, line, function, msg)
	}
}
*/
