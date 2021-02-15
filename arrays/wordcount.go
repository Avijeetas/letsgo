package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func prompt() int {

	fmt.Printf("===\t\tWord Counter\t\t===\n")
	fmt.Printf("\n1.	Take from user\n2.	Take from file\n3.	Exit\n\nEnter your choice: ")
	var typ int
	fmt.Scanln(&typ)
	return typ
}
func ReadUser() []string {
	arr := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	{
		fmt.Print("Enter the Text : ")

		scanner.Scan()

		text := scanner.Text()
		arr = append(arr, text)
	}
	return arr
}
func ReadFile() []string {
	var liness []string
	var s string
	fmt.Print("Enter the file name with correction extension : ")
	fmt.Scan(&s)
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		liness = append(liness, line)
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return liness
}
func done() {

	fmt.Printf("===\t\tThank you!\t\t===\n")
	os.Exit(0)
}
func main() {
	arrv := make([]string, 0)
	var typ int
	for {
		typ = prompt()
		if typ == 1 {
			arrv = ReadUser()

		} else if typ == 2 {
			arrv = ReadFile()
		} else {
			break
		}

		for _, val := range arrv {
			arr := counting(val)

			fmt.Printf("=================================================================\n")
			fmt.Printf("Number of characters : %d\n", arr[2])
			fmt.Printf("Number of words: %d\nNumber of unique words: %d\nNumber of Unique letters : %d", arr[1], arr[3], arr[0])
			fmt.Printf("\n")
		}

	}
	done()

}

func counting(str string) []int {
	count := make(map[rune]int)
	arr := make([]int, 4)
	word := 0
	ok := 0
	str += " "
	var wrds string
	//	l := len(str)
	//	fmt.Println(str)
	words := make(map[string]bool)
	for _, val := range str {
		//	fmt.Println(val)
		if string(val) == " " && ok == 1 {
			if words[wrds] == false {
				words[wrds] = true
				wrds = ""
			}
			word += 1
			ok ^= 1
		} else {

			ok = 1
			wrds += string(val)
			count[rune(val)]++
		}

	}

	arr[0] = len(count)
	arr[1] = word
	arr[2] = len(str) - 1
	arr[3] = len(words)
	return arr
}
