package main

import (
	"fmt"
	"github.com/mailgun/godebug/lib"
)

var func_lit_in_goScope = godebug.EnteringNewScope()

func main() {
	godebug.Line()
	hi, there := foo(7, 12)
	godebugScope := func_lit_in_goScope.EnteringNewChildScope()
	defer godebugScope.End()
	godebugScope.Declare("hi", &hi, "there", &there)
	godebug.Line()
	fmt.Println(hi, there)
}

var foo = func(a, _ int) (b, _ string) {
	var godebugResult2 string
	godebug.EnterFunc(func() {
		b, godebugResult2 = func() (b, _ string) {
			godebugScope := func_lit_in_goScope.EnteringNewChildScope()
			defer godebugScope.End()
			godebugScope.Declare("a", &a, "b", &b)
			godebug.Line()
			return "Hello", "World"
		}()
	})
	godebug.ExitFunc()
	return b, godebugResult2
}
