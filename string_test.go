package validator

import (
	"testing"
)

type testStringModel struct {
	Url   string `is:"url"`
	Email string `is:"email"`
	Min   string `min:"10"`
	Max   string `max:"20"`
}

func testString() {
	hand := new(Handler)
	hand.Attach(&String{})

	var ts *testStringModel
	ts = new(testStringModel)
	ts.Email = "test@email"
	ts.Url = "test.com"
	ts.Min = "test"
	ts.Max = "test.test.test.test"
	hand.Validate(ts, nil)
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testString()
	}
}

func TestString(t *testing.T) {

	hand := new(Handler)
	hand.Attach(&String{})

	var ts *testStringModel

	ts = new(testStringModel)
	ts.Email = "test@email"

	if errs := hand.Validate(ts, []string{"Email"}); errs == nil {
		t.Error("is email does not work")
	}

	ts = new(testStringModel)
	ts.Email = "test@email.com"
	if errs := hand.Validate(ts, []string{"Email"}); errs != nil {
		t.Error("is email does not work")
	}

	ts = new(testStringModel)
	ts.Url = "test"
	if errs := hand.Validate(ts, []string{"Url"}); errs == nil {
		t.Error("is url does not work")
	}

	ts = new(testStringModel)
	ts.Url = "test.com"
	if errs := hand.Validate(ts, []string{"Url"}); errs != nil {
		t.Error("is url does not work")
	}

	ts = new(testStringModel)
	ts.Url = "http://www.test.com?p1=1&p2=2"
	if errs := hand.Validate(ts, []string{"Url"}); errs != nil {
		t.Error("is url does not work")
	}

	ts = new(testStringModel)
	ts.Min = "test"
	if errs := hand.Validate(ts, []string{"Min"}); errs == nil {
		t.Error("min does not work")
	}

	ts = new(testStringModel)
	ts.Min = "test test test test test test"
	if errs := hand.Validate(ts, []string{"Min"}); errs != nil {
		t.Error("min does not work")
	}

	ts = new(testStringModel)
	ts.Max = "test test"
	if errs := hand.Validate(ts, []string{"Max"}); errs != nil {
		t.Error("max does not work")
	}

	ts = new(testStringModel)
	ts.Max = "test test test test test test"
	if errs := hand.Validate(ts, []string{"Max"}); errs == nil {
		t.Error("max does not work")
	}
}
