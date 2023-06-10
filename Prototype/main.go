package main

import (
	"fmt"
)

type interfacePrototype interface {
	clone() interfacePrototype
}

type files struct {
	name string
	typ  string
}

func (f *files) clone() interfacePrototype {
	return &files{
		name: f.name + "_clone",
		typ:  f.typ,
	}
}

// type folders struct {
// }

// func (f *folders) clone() {
// 	// do nothing
// }

func main() {
	f := files{
		name: "file1",
		typ:  "txt",
	}

	c := f.clone()

	fmt.Println(c)
}
