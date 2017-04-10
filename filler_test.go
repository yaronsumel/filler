package filler

import (
	"errors"
	"testing"
)

//go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html

var demoFiller = filler{
	tag: "demoFiller1",
	fn: func(obj interface{}) (interface{}, error) {
		return "hello", nil
	},
}

type demoStruct struct {
	Name    string `fill:"demoFiller1:Val"`
	Val     string `fill:"demoFiller2"`
	Ptr     *string
	XPtr    *string `fill:"fillPtr:Ptr"`
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
	RegFiller("demoFiller1", func(obj interface{}) (interface{}, error) {
		return "hello", nil
	})
	v1, err1 := fillers[0].fn("hello")
	v2, err2 := demoFiller.fn("hello")
	if fillers[0].tag != demoFiller.tag || v1 != v2 || err1 != err2 {
		t.FailNow()
	}
}

// Fill - fill the object with all the current fillers
func TestFill(t *testing.T) {
	RegFiller("demoFiller1", func(obj interface{}) (interface{}, error) {
		return "hello", nil
	})
	RegFiller("demoFillerErr", func(obj interface{}) (interface{}, error) {
		return nil, errors.New("some error")
	})
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
