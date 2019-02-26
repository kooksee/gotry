package gotry

import (
	"fmt"
	"github.com/juju/errors"
	"reflect"
	"strings"
)

func assertFn(fn interface{}) {
	assert(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	assert(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())
}

func assert(b bool, text string, args ...interface{}) {
	if b {
		panic(fmt.Sprintf(text, args...))
	}
}

type Assert struct {
}

func (t *Assert) Assert(b bool, format string, args ...interface{}) {
	if b {
		err := errors.NewErr(format, args...)
		err.SetLocation(1)
		panic(strings.Join(err.StackTrace(), "\n"))
	}
}

func (t *Assert) AssertErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	e := errors.Annotatef(err, format, args...)
	e.(*errors.Err).SetLocation(1)
	panic(errors.ErrorStack(e))
}

func (t *Assert) MustNotError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
