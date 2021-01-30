// slice is dynamic
// unlike arrays, funcations are available that allow us to add extra elements onto
//	the end of a slice;

package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7) // make(type, length)

	notes[0] = "f"
	notes[1] = "o"
	notes[2] = "o"
	notes[3] = "b"
	notes[4] = "a"
	notes[5] = "r"
	mySlice := notes[3:6]
	fmt.Println(notes)
	fmt.Println(mySlice)
	mySlice = append(mySlice, notes[0], notes[1], notes[2])
	//fmt.Println(mySlice)
	for _, val := range mySlice {
		fmt.Println(val)
	}
}
