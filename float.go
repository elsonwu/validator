package validator

import (
	"errors"
	"reflect"
	"strconv"
)

type Float struct{}

func (self *Float) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Float32 || f.Type.Kind() == reflect.Float64
}

func (self *Float) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && 0 == fv.Float() {
		errs = append(errs, errors.New(f.Name+" cannot be blank"))
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.ParseFloat(min, 64)
		if err != nil {
			errs = append(errs, err)
		}

		if min2 > fv.Float() {
			errs = append(errs, errors.New(f.Name+": min err"))
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.ParseFloat(max, 64)
		if err != nil {
			errs = append(errs, err)
		}

		if max2 < fv.Float() {
			errs = append(errs, errors.New(f.Name+": max err"))
		}
	}

	return
}