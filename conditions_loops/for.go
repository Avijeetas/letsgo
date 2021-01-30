package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input int
	fmt.Print("Enter a number more than 6 :")
	fmt.Scan(&input)

	for i := 1; i < input; i++ {

		if i == 2 {
			fmt.Println("Skipping")
		}
		x := strconv.Itoa(i)
		fmt.Println("Now at " + x)
		if i == 6 {
			x := "Now at " + x + "\nBreaking the loop"
			fmt.Println(x)
			break
		}

	}
}
