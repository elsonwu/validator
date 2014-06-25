package validator

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/elsonwu/i18n"
)

type Float struct {
	CoreValidate
}

func (self *Float) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Float32 || f.Type.Kind() == reflect.Float64
}

func (self *Float) Validate(f reflect.StructField, fv reflect.Value) []error {
	if "required" == f.Tag.Get("required") && 0 == fv.Float() {
		return []error{errors.New(i18n.T("%s cannot be blank", FieldName(f)))}
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.ParseFloat(min, 64)
		if err != nil {
			return []error{err}
		}

		if min2 > fv.Float() {
			return []error{errors.New(i18n.T("%s is too small", FieldName(f)))}
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.ParseFloat(max, 64)
		if err != nil {
			return []error{err}
		}

		if max2 < fv.Float() {
			return []error{errors.New(i18n.T("%s is too small", FieldName(f)))}
		}
	}

	return nil
}
