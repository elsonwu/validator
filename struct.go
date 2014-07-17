package validator

import (
	"reflect"

	"github.com/elsonwu/is"
)

type Struct struct {
	CoreValidate
}

func (self *Struct) Filter(f reflect.StructField, fv reflect.Value) bool {
	if reflect.Struct != fv.Kind() {
		return false
	}

	if "omitempty" == f.Tag.Get("omitempty") && is.Zero(fv) {
		return false
	}

	return true
}

func (self *Struct) Validate(f reflect.StructField, fv reflect.Value) []error {
	if fv.CanAddr() {
		if es := self.CoreValidate.handler.Validate(fv.Addr().Interface(), nil); es != nil {
			return es
		}
	}

	return nil
}
