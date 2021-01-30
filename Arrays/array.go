package main

import "fmt"

func main() {

	arr := [7]int{1, 2, 3, 5, 6, 7}
	fmt.Println(len(arr))

	user input
	var n int
	fmt.Scan(&n)
	notes := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&notes[i])
	}
	fmt.Println(notes)

}
