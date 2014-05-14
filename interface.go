package validator

import (
	"errors"
	"reflect"
)

type Interface struct {
	CoreValidate
}

func (self *Interface) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Interface
}

func (self *Interface) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		errs = append(errs, errors.New(f.Name+" cannot be blank"))
	}

	return
}
