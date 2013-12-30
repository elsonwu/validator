package validator

import (
	"errors"
	"reflect"
	"strconv"
)

type Int struct{}

func (self *Int) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Int ||
		f.Type.Kind() == reflect.Int8 ||
		f.Type.Kind() == reflect.Int16 ||
		f.Type.Kind() == reflect.Int32 ||
		f.Type.Kind() == reflect.Int64
}

func (self *Int) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			errs = append(errs, err)
		}

		if int64(min2) > fv.Int() {
			errs = append(errs, errors.New(f.Name+": min err"))
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			errs = append(errs, err)
		}

		if int64(max2) < fv.Int() {
			errs = append(errs, errors.New(f.Name+": max err"))
		}
	}

	return
}
