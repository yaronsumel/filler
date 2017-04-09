# filler [![Go Report Card](https://goreportcard.com/badge/github.com/yaronsumel/filler)](https://goreportcard.com/report/github.com/yaronsumel/filler) [![Build Status](https://travis-ci.org/yaronsumel/grapes.svg?branch=master)](https://travis-ci.org/yaronsumel/filler) [![GoDoc](https://godoc.org/github.com/yaronsumel/filler?status.svg)](https://godoc.org/github.com/yaronsumel/filler)
###### small util to fill gaps in your structs 

Installation
------
```bash
$ go get github.com/yaronsumel/filler
```

[Usage](https://github.com/yaronsumel/filler/blob/master/example/example.go)
------

```go
package main

import (
	"fmt"
	"github.com/yaronsumel/filler"
)

type model struct {
	UserID   string
	// will pass the UserID val into UserNameFiller Fn
	UserName string `fill:"UserNameFiller:UserID"`
}

func init() {
	// register the filler
	filler.RegFiller(filler.Filler{
		Tag: "UserNameFiller",
		Fn: func(value interface{}) (interface{}, error) {
			return "UserId" + value.(string), nil
		},
	})
}

func main() {
	m := &model{
		UserID: "123",
	}
	fmt.Printf("%+v\n", m)
	// should print `&{UserId:123 UserName:}`
	filler.Fill(m)
	// should print `&{UserId:123 UserName:UserId123}`
	fmt.Printf("%+v\n", m)
}

```

> ##### Written and Maintained by [@YaronSumel](https://twitter.com/yaronsumel) #####
