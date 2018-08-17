package main

import (
	"fmt"
	"github.com/mozey/go-cookbook/ldflags/config"
)

var foo string

func main() {
	// Private package var
	if foo == "" {
		foo = "bla"
	}
	// Public var in imported package
	if config.Foo == "" {
		config.Foo = "bla"
	}
	// Private var in imported package
	fiz := config.GetFiz()
	fmt.Println(fmt.Sprintf("%v %v %v", foo, config.Foo, fiz))
}