package gotry

import (
	"errors"
	"reflect"
)

type _try struct {
	err    error
	params []reflect.Value
}

func (t *_try) Catch(fn func(err error)) {

	if t.err == nil {
		return
	}

	fn(t.err)
}

func Try(fn func(assert *Assert)) *_try {
	assertFn(fn)

	t := &_try{}
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				case error:
					t.err = r.(error)
				case string:
					t.err = errors.New(r.(string))
				}
			}
		}()
		t.params = reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(new(Assert))})
	}()
	return t
}
