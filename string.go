package validator

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/elsonwu/i18n"
	"github.com/elsonwu/is"
)

type String struct {
	CoreValidate
}

func (self *String) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.String
}

func (self *String) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && "" == fv.String() {
		errs = append(errs, errors.New(i18n.T("%s cannot be blank", FieldName(f))))
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			errs = append(errs, err)
		}

		if min2 > fv.Len() {
			errs = append(errs, errors.New(i18n.T("%s min err", FieldName(f))))
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			errs = append(errs, err)
		}

		if max2 < fv.Len() {
			errs = append(errs, errors.New(i18n.T("%s max err", FieldName(f))))
		}
	}

	isv := f.Tag.Get("is")
	if "" != isv {
		switch isv {
		case "email":
			if !is.Email(fv.String()) {
				errs = append(errs, errors.New(i18n.T("%s is not a valid email", FieldName(f))))
			}
		case "url":
			if !is.Url(fv.String()) {
				errs = append(errs, errors.New(i18n.T("%s is not a valid url", FieldName(f))))
			}
		}
	}

	return
}
