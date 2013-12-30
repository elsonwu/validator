package validator

import (
	"reflect"
)

type IValidator interface {
	Filter(f reflect.StructField, fv reflect.Value) bool
	Validate(f reflect.StructField, fv reflect.Value) (errs []error)
}

type IValidatorHandler interface {
	validateField(f reflect.StructField, fv reflect.Value) (errs []error)
	Validate(m interface{}, attributes []string) (errs []error)
	Attach(v IValidator)
}
type IValidateModel interface {
	Validate() []error
}
