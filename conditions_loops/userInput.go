package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// err is declared but no used
	//	input, err := reader.ReadString('\n')
	//_ is a blank indicator and used for ignoring undeclared variable

	//input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	// String to integer conversion

	input = strings.TrimSpace(input)

	grade, err := strconv.Atoi(input)
	fmt.Println(reflect.TypeOf(input))
	fmt.Println(grade)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(grade))
	fmt.Println("Input is converted into integer")
}
