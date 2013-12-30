package validator

import (
	"testing"
)

type testFloatModel struct {
	Min float32 `min:"10.11"`
	Max float32 `max:"20.11"`
}

func TestFloat(t *testing.T) {

	hand := new(Handler)
	hand.Attach(&Float{})

	var ts *testFloatModel

	ts = new(testFloatModel)
	ts.Min = float32(1.1)
	if errs := hand.Validate(ts, []string{"Min"}); errs == nil {
		t.Error("min does not work")
	}

	ts = new(testFloatModel)
	ts.Min = float32(10.12)
	if errs := hand.Validate(ts, []string{"Min"}); errs != nil {
		t.Error("min does not work")
	}

	ts = new(testFloatModel)
	ts.Min = float32(12.0)
	if errs := hand.Validate(ts, []string{"Min"}); errs != nil {
		t.Error("min does not work")
	}

	ts = new(testFloatModel)
	ts.Max = float32(1.1)
	if errs := hand.Validate(ts, []string{"Max"}); errs != nil {
		t.Error("max does not work")
	}

	ts = new(testFloatModel)
	ts.Max = float32(20.0)
	if errs := hand.Validate(ts, []string{"Max"}); errs != nil {
		t.Error("max does not work")
	}

	ts = new(testFloatModel)
	ts.Max = float32(23.1)
	if errs := hand.Validate(ts, []string{"Max"}); errs == nil {
		t.Error("max does not work")
	}
}
