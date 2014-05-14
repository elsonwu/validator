package validator

import (
	"errors"
	"reflect"
	"strings"
)

var defaultHandler IValidatorHandler

func DefaultHandler() IValidatorHandler {
	if defaultHandler == nil {
		defaultHandler = new(Handler)
		defaultHandler.Attach(&Array{})
		defaultHandler.Attach(&Int{})
		defaultHandler.Attach(&Float{})
		defaultHandler.Attach(&String{})
		defaultHandler.Attach(&Map{})
		defaultHandler.Attach(&Interface{})
	}

	return defaultHandler
}

type Handler struct {
	validators []IValidator
}

func (self *Handler) Attach(v IValidator) {
	v.SetHandler(self)
	self.validators = append(self.validators, v)
}

func (self *Handler) ValidateField(f reflect.StructField, fv reflect.Value) (errs []error) {

	if reflect.Struct == f.Type.Kind() {
		if es := self.Validate(fv.Interface(), nil); es != nil {
			errs = append(errs, es...)
		}
	}

	for _, v := range self.validators {
		if v.Filter(f, fv) {
			if es := v.Validate(f, fv); es != nil {
				errs = append(errs, es...)
			}
		}
	}

	return
}

func (self *Handler) Validate(m interface{}, attributes []string) (errs []error) {

	refType := reflect.TypeOf(m)
	refValue := reflect.ValueOf(m)
	if refType.Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	if refType.Kind() != reflect.Struct {
		return []error{errors.New("The validate m must be struct or ptr")}
	}

	if attributes != nil {
		for _, name := range attributes {
			//Handle embedded attributes, for example: Profile.SecondaryEmail or event Comments.Id
			if _names := strings.SplitAfterN(name, ".", 2); 1 < len(_names) {
				attr := strings.Trim(_names[0], ".")
				if attrT, ok := refType.FieldByName(attr); ok {
					if attrT.Type.Kind() != reflect.Struct {
						errs = append(errs, errors.New(attr+" is not a struct"))
						return
					}

					if fv := refValue.FieldByName(attr); fv.IsValid() {
						if _errs := self.Validate(fv.Interface(), []string{_names[1]}); _errs != nil {
							errs = append(errs, _errs...)
						}
					} else {
						errs = append(errs, errors.New(attr+" is not valid"))
					}
				} else {
					errs = append(errs, errors.New(attr+" does not exists"))
				}
			} else { //no ".", so we validate directly
				if f, ok := refType.FieldByName(name); ok {
					if fv := refValue.FieldByName(name); fv.IsValid() {
						if es := self.ValidateField(f, fv); es != nil {
							errs = append(errs, es...)
						}
					} else {
						errs = append(errs, errors.New(name+" does not exists"))
					}
				} else {
					errs = append(errs, errors.New(name+" does not exists"))
				}
			}
		}
	} else {
		for i, numFile := 0, refType.NumField(); i < numFile; i++ {
			f := refType.Field(i)
			fv := refValue.Field(i)
			if es := self.ValidateField(f, fv); es != nil {
				errs = append(errs, es...)
			}
		}
	}

	if im, ok := m.(IValidateModel); ok {
		if es := im.Validate(); es != nil {
			errs = append(errs, es...)
		}
	}

	return
}
