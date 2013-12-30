A simple struct validator
=========================

	type Model struct {
    	A int `min:"10" max:"100"`
	    B float `min:"1.1" max:"3.3"`
    	C array `min:"1" max:"3"`
	    D map `keys:"k1|k2|k3"`
    	E string `min:"10" max:"100"`
	    Email string `is:"email"`
    	Url string `is:"url"`
	}

	model := new(Model)
	validatorHandler := validator.DefaultHandler()
	
	//validate the whold model
	errs := validatorHandler.Validate(model, nil)
	if errs != nil {
		//...
	}
	
	//validate the given attributes.
	errs = validatorHandler.Validate(model, []string{"A", "Email"})
	if errs != nil {
		//...
	}

#customized validators

	//you can register the validator by yourself
	validatorHandler := new(validator.Handler)
	
	//use any validator as you want
	validatorHandler.Attach(&validator.String{})
	validatorHandler.Attach(&YourValidator{})
	
#user-defined validator

	type YourValidator struct {}
	func (self *YourValidator) Filter(f reflect.StructField, fv reflect.Value) bool {
		//filter the type of field you want to validate
		return true	
	}
	
	func (self *Int) Validate(f reflect.StructField, fv reflect.Value) (errs []error) {
	
		//then this validator will handle the "test" tag
		test := f.Tag.Get("test")
		
		//validate the test with the field's value
		
		return
	}