package validator

import (
	"errors"
	"reflect"
	"strconv"
)

type Array struct {
	CoreValidate
}

func (self *Array) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Array || f.Type.Kind() == reflect.Slice
}

func (self *Array) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		errs = append(errs, errors.New(f.Name+" cannot be blank"))
	}

	// if the item in array/slice is struct,
	// we need to run handler.Validate again.
	if !fv.IsNil() {
		for i := 0; i < fv.Len(); i++ {
			if v := fv.Index(i); v.Kind() == reflect.Struct {
				errs = append(errs, self.CoreValidate.handler.Validate(v.Interface(), nil)...)
			}
		}
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			errs = append(errs, err)
		}

		if min2 > fv.Len() {
			errs = append(errs, errors.New(f.Name+": min err"))
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			errs = append(errs, err)
		}

		if max2 < fv.Len() {
			errs = append(errs, errors.New(f.Name+": max err"))
		}
	}

	return
}
