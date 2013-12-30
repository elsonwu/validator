package validator

import (
	"reflect"
)

type IValidator interface {
	Filter(f reflect.StructField, fv reflect.Value) bool
	Validate(f reflect.StructField, fv reflect.Value) (errs []error)
}

type IValidateModel interface {
	Validate() []error
}
