package main

import (
	"github.com/yaronsumel/filler"
	"fmt"
)

type Model struct {
	UserId   string
	UserName string        `fill:"UserNameFiller:UserId"`
}

func init(){
	filler.RegFiller(filler.Filler{
		Tag: "UserNameFiller",
		Fn: func(value interface{}) (interface{}, error) {
			return "UserId"+value.(string), nil
		},
	})
}

func main() {
	m := &Model{
		UserId:"123",
	}
	fmt.Printf("%+v\n", m)
	// should print `&{UserId:123 UserName:}`
	filler.Fill(m)
	// should print `&{UserId:123 UserName:UserId123}`
	fmt.Printf("%+v\n", m)
}