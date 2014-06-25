package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/elsonwu/i18n"
)

type Map struct {
	CoreValidate
}

func (self *Map) Filter(f reflect.StructField, fv reflect.Value) bool {
	return f.Type.Kind() == reflect.Map
}

func (self *Map) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	if "required" == f.Tag.Get("required") && fv.IsNil() {
		errs = append(errs, errors.New(i18n.T("%s cannot be blank", FieldName(f))))
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
		errs = append(errs, errors.New(i18n.T("%s keys cannot be blank", FieldName(f))))
	}

	if len(ks) != len(keys2) {
		errs = append(errs, errors.New(i18n.T("%s count of keys is not the same as setting", FieldName(f))))
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
				errs = append(errs, errors.New(i18n.T("%s key %s does not exists", FieldName(f), k)))
			}
		}
	}

	return
}
