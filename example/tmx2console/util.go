package main

import (
	"errors"
	"github.com/ChaimHong/go-tmx"
)

var (
	PropertyNotUnique   = errors.New("Layer Property is not unique")
	PropertyUnavailable = errors.New("Property does not exist")
)

func GetProperty(p *[]tmx.Property, name string) (value string, err error) {
	for _, p0 := range *p{
		if p0.Name ==  name{
			return p0.Value, nil
		}
	}

	err = PropertyUnavailable
	return
}
