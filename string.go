package validator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

const (
	EmailPattern = `^[a-zA-Z0-9!#$%&\'*+\\/=?^_{|}~-]+(?:\.[a-zA-Z0-9!#$%&\'*+\\/=?^_{|}~-]+)*@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`
	UrlPattern   = `((http|https|ftp){1}\:\/\/)?([a-zA-Z0-9\-\.]+[a-zA-Z0-9]+\.(net|cn|co|hk|tw|com|edu|gov|us|int|mil|org|int|mil|vg|uk|idv|tk|se|nz|nu|nl|ms|jp|jobs|it|ind|gen|firm|in|gs|fr|fm|eu|es|de|bz|be|at|am|ag|mx|asia|ws|xxx|tv|cc|ca|mobi|me|biz|arpa|info|name|pro|aero|coop|museum|ly|eg|mk)(:[a-zA-Z0-9]*)?\/?([a-zA-Z0-9\-\._\?\'\/\\\+&amp;%\$#\=~])*)+`
)

type String struct {
}

func (self *String) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.String
}

func (self *String) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && "" == fv.String() {
		errs = append(errs, errors.New(f.Name+" cannot be blank"))
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

	is := f.Tag.Get("is")
	if "" != is {
		switch is {
		case "email":
			if ok, err := regexp.MatchString(EmailPattern, fv.String()); !ok {
				errs = append(errs, errors.New(f.Name+": is not a valid email"))
			} else if err != nil {
				errs = append(errs, err)
			}
		case "url":
			if ok, err := regexp.MatchString(UrlPattern, fv.String()); !ok {
				errs = append(errs, errors.New(f.Name+": is not a valid url"))
			} else if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return
}
