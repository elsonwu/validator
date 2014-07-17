package validator

import (
	"errors"
	"reflect"

	"github.com/elsonwu/i18n"
)

type Interface struct {
	CoreValidate
}

func (self *Interface) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Interface
}

func (self *Interface) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		errs = append(errs, errors.New(i18n.T("%s cannot be blank", f.Name)))
	}

	if !fv.IsNil() && fv.Elem().Kind() == reflect.Ptr {
		if errs := self.CoreValidate.handler.Validate(fv.Elem().Interface(), nil); errs != nil {
			return errs
		}
	}

	return
}
