package validator

import (
	"errors"
	"reflect"
	"strings"
)

type Map struct{}

func (self *Map) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Map
}

func (self *Map) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		errs = append(errs, errors.New(f.Name+" cannot be blank"))
	}

	ruleKeys := f.Tag.Get("keys")
	if "" == ruleKeys {
		return
	}

	ks := strings.Split(ruleKeys, "|")
	if nil == ks {
		return
	}

	keys2 := fv.MapKeys()
	if nil == keys2 {
		errs = append(errs, errors.New(f.Name+": keys cannot be blank"))
	}

	if len(ks) != len(keys2) {
		errs = append(errs, errors.New(f.Name+": count of keys is not the same as setting"))
	} else {
		var find bool
		for _, k := range ks {
			find = false
			for _, k2 := range keys2 {
				if k == k2.String() {
					find = true
					break
				}
			}

			if !find {
				errs = append(errs, errors.New(f.Name+": key("+k+") does not exists"))
			}
		}
	}

	return
}
