package main

import (
	"fmt"
)

type interfacePrototype interface {
	clone() interfacePrototype
}

type files struct {
	name         string
	fileType     string
	architecture string
}

func (f *files) clone() interfacePrototype {
	return &files{
		name:         f.name + "_clone",
		fileType:     f.fileType,
		architecture: f.architecture,
	}
}

// Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.
// https://refactoring.guru/design-patterns/prototype
// Personal note: It simply copies the object by having a clone method within the class itself rather than calling the constructor and copying from it.

func main() {
	f := files{
		name:         "file1",
		fileType:     "txt",
		architecture: "x64",
	}

	c := f.clone()
	fmt.Println(c)
}
