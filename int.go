package validator

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/elsonwu/i18n"
)

type Int struct {
	CoreValidate
}

func (self *Int) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Int ||
		f.Type.Kind() == reflect.Int8 ||
		f.Type.Kind() == reflect.Int16 ||
		f.Type.Kind() == reflect.Int32 ||
		f.Type.Kind() == reflect.Int64
}

func (self *Int) Validate(f reflect.StructField, fv reflect.Value) []error {
	if "required" == f.Tag.Get("required") && 0 == fv.Int() {
		return []error{errors.New(i18n.T("%s cannot be blank", FieldName(f)))}
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			return []error{err}
		}

		if int64(min2) > fv.Int() {
			return []error{errors.New(i18n.T("%s is too small", FieldName(f)))}
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			return []error{err}
		}

		if int64(max2) < fv.Int() {
			return []error{errors.New(i18n.T("%s is too big", FieldName(f)))}
		}
	}

	return nil
}
