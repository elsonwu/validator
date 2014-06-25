package validator

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/elsonwu/i18n"
)

type Array struct {
	CoreValidate
}

func (self *Array) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Array || f.Type.Kind() == reflect.Slice
}

func (self *Array) Validate(f reflect.StructField, fv reflect.Value) []error {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		return []error{errors.New(i18n.T("%s cannot be blank", FieldName(f)))}
	}

	// if the item in array/slice is struct,
	// we need to run handler.Validate again.
	if !fv.IsNil() {
		for i := 0; i < fv.Len(); i++ {
			if v := fv.Index(i); v.Kind() == reflect.Struct {
				if v.CanAddr() && v.Addr().CanInterface() {
					return self.CoreValidate.handler.Validate(v.Addr().Interface(), nil)
				}
			}
		}
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			return []error{err}
		}

		if min2 > fv.Len() {
			return []error{errors.New(i18n.T("%s's length is too small", FieldName(f)))}
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			return []error{err}
		}

		if max2 < fv.Len() {
			return []error{errors.New(i18n.T("%s's length  is too big", FieldName(f)))}
		}
	}

	return nil
}
