package validator

import (
	"reflect"
)

type CoreValidate struct {
	handler IValidatorHandler
}

func (self *CoreValidate) SetHandler(handler IValidatorHandler) {
	self.handler = handler
}

type IValidator interface {
	SetHandler(IValidatorHandler)
	Filter(f reflect.StructField, fv reflect.Value) bool
	Validate(f reflect.StructField, fv reflect.Value) (errs []error)
}

type IValidatorHandler interface {
	ValidateField(f reflect.StructField, fv reflect.Value) (errs []error)
	Validate(m interface{}, attributes []string) (errs []error)
	Attach(v IValidator)
}

type IValidateModel interface {
	Validate() []error
}

type IBeforeValidateModel interface {
	BeforeValidate() []error
}
