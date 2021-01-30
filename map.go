package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var liness []string
	file, err := os.Open("friends.txt")
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

	count := make(map[string]int)
	for _, i := range liness {
		count[i]++
	}

	for i, f := range count {
		fmt.Printf("user[ %s ] =%d\n", i, f)
	}
}
