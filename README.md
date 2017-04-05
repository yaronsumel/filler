# filler [![Go Report Card](https://goreportcard.com/badge/github.com/yaronsumel/filler)](https://goreportcard.com/report/github.com/yaronsumel/filler) [![GoDoc](https://godoc.org/github.com/yaronsumel/filler?status.svg)](https://godoc.org/github.com/yaronsumel/filler)
###### small util to fill gaps in your structs 

Installation
------
```bash
$ go get github.com/yaronsumel/filler
```

Usage
------

##### fill tag

###### `fill:"[FillerName:OptionalValue]"`
###### `fill:"[User:UserId]"` - Fill current filed with the "User" Filler and UserId value
###### `fill:"[SayHello]"` = Fill current with "SayHello" Filler Without any value 


###### Add the `fill` tag in your model
```go
type Model struct {
	UserId   bson.ObjectId `json:"userId" bson:"userId"`
	FieldA   string        `json:"FieldA" bson:"FieldA" fill:"SayHello"`
	UserName string        `json:"user" bson:"-" fill:"User:UserId"`
}
```
###### Register the fillers
```go
	filler.RegFiller(filler.Filler{
		Tag: "User",
		Fn: func(value interface{}) (interface{}, error) {
			return "this is the user name", nil
		},
	})

	filler.RegFiller(filler.Filler{
		Tag: "SayHello",
		Fn: func(value interface{}) (interface{}, error) {
			return "Hello", nil
		},
	})
```

###### and Fill
```go
	m := Model{}
	filler.Fill(&m)
```

TBD
------
* testing

> ##### Written and Maintained by [@YaronSumel](https://twitter.com/yaronsumel) #####
