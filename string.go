package validator

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/elsonwu/i18n"
	"github.com/elsonwu/is"
)

type String struct {
	CoreValidate
}

func (self *String) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.String
}

func (self *String) Validate(f reflect.StructField, fv reflect.Value) []error {
	if "required" == f.Tag.Get("required") && "" == fv.String() {
		return []error{errors.New(i18n.T("%s cannot be blank", FieldName(f)))}
	}

	min := f.Tag.Get("min")
	if "" != min {
		min2, err := strconv.Atoi(min)
		if err != nil {
			return []error{err}
		}

		if min2 > fv.Len() {
			return []error{errors.New(i18n.T("%s is too short", FieldName(f)))}
		}
	}

	max := f.Tag.Get("max")
	if "" != max {
		max2, err := strconv.Atoi(max)
		if err != nil {
			return []error{err}
		}

		if max2 < fv.Len() {
			return []error{errors.New(i18n.T("%s is too long", FieldName(f)))}
		}
	}

	isv := f.Tag.Get("is")
	if "" != isv {
		tags := strings.Split(isv, ",")

		if 1 < len(tags) {
			if "omitempty" == tags[1] && fv.String() == "" {
				goto skip
			}
		}

		switch tags[0] {
		case "email":
			if !is.Email(fv.String()) {
				return []error{errors.New(i18n.T("%s is not a valid email", FieldName(f)))}
			}
		case "url":
			if !is.Url(fv.String()) {
				return []error{errors.New(i18n.T("%s is not a valid url", FieldName(f)))}
			}
		}
	skip:
	}

	return nil
}
