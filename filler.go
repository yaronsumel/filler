package filler

import (
	"reflect"
	"strings"
	"log"
)

const (
	tagName   = "fill"
	ignoreTag = "-"
	emptyTag  = ""
)

var fillers []Filler

// Filler instance
type Filler struct {
	// Tag is the prefix inside fill tag ie. "fill:mytag"
	Tag string
	// Fn function to call - helps us to fill the gaps
	Fn func(obj interface{}) (interface{}, error)
}

// RegFiller - register new filler into []fillers
func RegFiller(f Filler) {
	fillers = append(fillers, f)
}

// Fill - fill the object with all the current fillers
func Fill(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		log.Panic("panic at [yaronsumel/filler] : obj kind passed to Fill should be Ptr")
	}
	v := reflect.TypeOf(obj).Elem()
	s := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		currentField := v.Field(i)
		tag := currentField.Tag.Get(tagName)
		if tag == emptyTag || tag == ignoreTag {
			continue
		}
		t, elm := parseTag(tag)
		for _, filter := range fillers {
			var elmValue interface{}
			if filter.Tag == t {
				if elm != "" {
					elmValue = s.FieldByName(elm).Interface()
				}
				res, err := filter.Fn(elmValue)
				if err != nil {
					continue
				}
				if s.FieldByName(currentField.Name).CanSet() {
					s.FieldByName(currentField.Name).Set(reflect.ValueOf(res))
				}
			}
		}
	}
}

// parseTag split the string by ":" and return two strings
func parseTag(tag string) (string, string) {
	x := strings.Split(tag, ":")
	if len(x) != 2 {
		return x[0], ""
	}
	return x[0], x[1]
}
