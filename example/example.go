package main

import (
	"fmt"
	"github.com/yaronsumel/filler"
)

type model struct {
	UserID   string
	UserName string `fill:"UserNameFiller:UserID"`
}

func init() {
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
