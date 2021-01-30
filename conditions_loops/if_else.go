package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter a grade: ")
	var grade int
	fmt.Scan(&grade)
	if grade >= 80 {
		fmt.Println("A+")
	} else if grade >= 70 {
		fmt.Println("A")
	} else if grade >= 60 {
		fmt.Println("A-")
	} else {
		fmt.Println("Fail!")
	}
}
