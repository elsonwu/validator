package validator

import (
	"testing"
)

type testIntModel struct {
	Min int `min:"10"`
	Max int `max:"20"`
}

func TestInt(t *testing.T) {

	hand := new(Handler)
	hand.Attach(&Int{})

	var ts *testIntModel

	ts = new(testIntModel)
	ts.Min = 1
	if errs := hand.Validate(ts, []string{"Min"}); errs == nil {
		t.Error("min does not work")
	}

	ts = new(testIntModel)
	ts.Min = 10
	if errs := hand.Validate(ts, []string{"Min"}); errs != nil {
		t.Error("min does not work")
	}

	ts = new(testIntModel)
	ts.Min = 12
	if errs := hand.Validate(ts, []string{"Min"}); errs != nil {
		t.Error("min does not work")
	}

	ts = new(testIntModel)
	ts.Max = 1
	if errs := hand.Validate(ts, []string{"Max"}); errs != nil {
		t.Error("max does not work")
	}

	ts = new(testIntModel)
	ts.Max = 20
	if errs := hand.Validate(ts, []string{"Max"}); errs != nil {
		t.Error("max does not work")
	}

	ts = new(testIntModel)
	ts.Max = 23
	if errs := hand.Validate(ts, []string{"Max"}); errs == nil {
		t.Error("max does not work")
	}
}
