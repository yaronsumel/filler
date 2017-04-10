package filler

import (
	"errors"
	"github.com/golang/groupcache/singleflight"
	"github.com/mitchellh/hashstructure"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	tagName   = "fill"
	ignoreTag = "-"
	emptyTag  = ""
)

var fillers []*filler
var mu sync.Mutex

// Filler instance
type filler struct {
	// Tag is the prefix inside fill tag ie. "fill:mytag"
	tag string
	// Fn function to call - helps us to fill the gaps
	fn func(obj interface{}) (interface{}, error)
	// duplicate function call suppression
	singleFlightGroup singleflight.Group
}

// RegFiller - register new filler into []fillers
func RegFiller(name string, fn func(obj interface{}) (interface{}, error)) {
	mu.Lock()
	fillers = append(fillers, &filler{
		name,
		fn,
		singleflight.Group{},
	})
	mu.Unlock()
}

// Fill - fill the object with all the current fillers
func Fill(obj interface{}) error {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("[yaronsumel/filler]: obj kind passed to Fill should be Ptr")
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
		for key, filler := range fillers {
			var elmValue interface{}
			if filler.tag == t {
				if elm != "" {
					elmValue = s.FieldByName(elm).Interface()
				}
				// if fill got called more than once - will get called once per fillerTag+value
				res, err := fillers[key].singleFlightGroup.Do(hash(filler.tag, elmValue), func() (interface{}, error) {
					return filler.fn(elmValue)
				})
				if err != nil {
					return err
				}
				resVal := reflect.ValueOf(res)
				// return err if can not Set or not same Kind
				if !s.FieldByName(currentField.Name).CanSet() || resVal.Kind() != s.FieldByName(currentField.Name).Kind() {
					return errors.New("[yaronsumel/filler]: Could not set value from Fn")
				}
				s.FieldByName(currentField.Name).Set(resVal)
			}
		}
	}
	return nil
}

func hash(name string, value interface{}) string {
	hash, err := hashstructure.Hash(value, nil)
	if err != nil {
		return name + strconv.FormatInt(int64(time.Now().Nanosecond()), 10)
	}
	return name + strconv.FormatUint(hash, 10)
}

// parseTag split the string by ":" and return two strings
func parseTag(tag string) (string, string) {
	x := strings.Split(tag, ":")
	if len(x) != 2 {
		return x[0], ""
	}
	return x[0], x[1]
}
