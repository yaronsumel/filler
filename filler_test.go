package filler

import (
	"errors"
	"testing"
)

var demoFiller = Filler{
	Tag: "demoFiller1",
	Fn: func(obj interface{}) (interface{}, error) {
		return "hello", nil
	},
}

var errDemoFiller = Filler{
	Tag: "demoFillerErr",
	Fn: func(obj interface{}) (interface{}, error) {
		return nil, errors.New("some error")
	},
}

type demoStruct struct {
	Name    string `fill:"demoFiller1:Val"`
	Val     string `fill:"demoFiller2"`
	Ignore1 string `fill:"-"`
	Ignore2 string `fill:""`
}

type notSameTypeStruct struct {
	Val int `fill:"demoFiller1"`
}

type errFromFn struct {
	Val int `fill:"demoFillerErr"`
}

// RegFiller - register new filler into []fillers
func TestRegFiller(t *testing.T) {
	RegFiller(demoFiller)
	v1, err1 := fillers[0].Fn("hello")
	v2, err2 := demoFiller.Fn("hello")
	if fillers[0].Tag != demoFiller.Tag || v1 != v2 || err1 != err2 {
		t.FailNow()
	}
}

// Fill - fill the object with all the current fillers
func TestFill(t *testing.T) {
	RegFiller(demoFiller)
	RegFiller(errDemoFiller)
	m := demoStruct{
		Name: "nameVal",
		Val:  "valVal",
	}
	// check non ptr - should return error
	if err := Fill(m); err == nil {
		t.FailNow()
	}
	// check if got filled
	Fill(&m)
	// should be filled
	if m.Name != "hello" || m.Val != "valVal" {
		t.FailNow()
	}
	m2 := notSameTypeStruct{}
	if err := Fill(&m2); err == nil {
		t.FailNow()
	}
	m3 := errFromFn{}
	if err := Fill(&m3); err == nil {
		t.FailNow()
	}
}
