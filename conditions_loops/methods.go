/*
Methods are functions that are associated with values of a particular type.
functions that are associated with values of a given type. Go methods are
kind of like the methods that you may have seen attached to “objects” in other languages, but they’re a bit simpler.


for example.
fmt.Println()
here Println is a method*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# rocks!"

	// NewReplacer is used to set up to replace every # with o
	// Replace is a method name and replacer is a value
	replacer := strings.NewReplacer("#", "o")

	fixed := replacer.Replace(broken)

	fmt.Println(fixed)
}
