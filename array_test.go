package validator

import (
	"testing"
)

type testArrayModel struct {
	Att []string `min:"3" max:"10" required:"required"`
}

func TestArray1(t *testing.T) {
	tam := new(testArrayModel)
	tam.Att = []string{
		"test1",
		"test2",
		"test3",
		"test4",
	}
	hand := new(Handler)
	hand.Attach(&Array{})

	errs := hand.Validate(tam, nil)
	if errs != nil {
		t.Error("the rule min/max does not work")
	}
}

func TestArray2(t *testing.T) {
	tam := new(testArrayModel)
	hand := new(Handler)
	hand.Attach(&Array{})

	errs := hand.Validate(tam, nil)
	if errs == nil {
		t.Error("the rule min/max does not work")
	} else {
		var minFind, maxFind bool
		for _, err := range errs {
			if err.Error() == "Att: min err" {
				minFind = true
			}

			if err.Error() == "Att: max err" {
				maxFind = true
			}
		}

		if !minFind {
			t.Error("rule min does not work")
		}

		if maxFind {
			t.Error("rule max does not work")
		}
	}
}

func TestArray3(t *testing.T) {
	tam := new(testArrayModel)
	tam.Att = []string{"test1", "test2"}
	hand := new(Handler)
	hand.Attach(&Array{})

	errs := hand.Validate(tam, nil)
	if errs == nil {
		t.Error("the rule min/max does not work")
	} else {
		var minFind, maxFind bool
		for _, err := range errs {
			if err.Error() == "Att: min err" {
				minFind = true
			}

			if err.Error() == "Att: max err" {
				maxFind = true
			}
		}

		if !minFind {
			t.Error("rule min does not work")
		}

		if maxFind {
			t.Error("rule max does not work")
		}
	}
}

func TestArray4(t *testing.T) {
	tam := new(testArrayModel)
	tam.Att = []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
		"test6",
		"test7",
		"test8",
		"test9",
		"test10",
		"test11",
	}
	hand := new(Handler)
	hand.Attach(&Array{})

	errs := hand.Validate(tam, nil)
	if errs == nil {
		t.Error("the rule min/max does not work")
	} else {
		var minFind, maxFind bool
		for _, err := range errs {
			if err.Error() == "Att: min err" {
				minFind = true
			}

			if err.Error() == "Att: max err" {
				maxFind = true
			}
		}

		if minFind {
			t.Error("rule min does not work")
		}

		if !maxFind {
			t.Error("rule max does not work")
		}
	}
}

func TestArray5(t *testing.T) {
	tam := new(testArrayModel)
	hand := new(Handler)
	hand.Attach(&Array{})
	errs := hand.Validate(tam, nil)
	if errs == nil {
		t.Error("the rule does not work")
	} else {
		var find bool
		for _, err := range errs {
			if err.Error() == "Att cannot be blank" {
				find = true
			}
		}

		if !find {
			t.Error("the rule required does not work")
		}
	}
}
