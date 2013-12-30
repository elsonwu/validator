package validator

import (
	"testing"
)

type testMapModel struct {
	Map map[string]int `keys:"test1|test2"`
}

func TestMap(t *testing.T) {
	hand := new(Handler)
	hand.Attach(&Map{})

	var tm *testMapModel

	tm = new(testMapModel)
	tm.Map = map[string]int{"test2": 2}

	if errs := hand.Validate(tm, []string{"Map"}); errs == nil {
		t.Error("map keys does not work")
	}

	tm = new(testMapModel)
	tm.Map = map[string]int{"test": 1, "test1": 2}

	if errs := hand.Validate(tm, []string{"Map"}); errs == nil {
		t.Error("map keys does not work")
	}

	tm = new(testMapModel)
	tm.Map = map[string]int{"test1": 1, "test2": 2}

	if errs := hand.Validate(tm, []string{"Map"}); errs != nil {
		t.Error("map keys does not work")
	}

	tm = new(testMapModel)
	tm.Map = map[string]int{"test1": 1, "test2": 2, "test3": 3}

	if errs := hand.Validate(tm, []string{"Map"}); errs == nil {
		t.Error("map keys does not work")
	}

	tm = new(testMapModel)
	tm.Map = map[string]int{"test1": 1, "test2": 2}

	if errs := hand.Validate(tm, []string{"Map"}); errs != nil {
		t.Error("map keys does not work")
	}
}
