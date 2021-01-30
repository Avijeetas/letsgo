package main

import "fmt"

func area(width float64, length float64) {
	totalArea := width * length

	fmt.Printf("%s is %.2f\n", "Area", totalArea)

}

func inp() (float64, float64) {
	var length float64
	var width float64

	fmt.Scan(&length)
	fmt.Scan(&width)
	return length, width
}
func main() {
	length, width := inp()
	fmt.Printf("Length is %.2f Width is %.2f\n", length, width)

	area(width, length)

}
