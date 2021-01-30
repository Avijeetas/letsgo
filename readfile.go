package readfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getinp() {
	var liness []string
	file, err := os.Open("friends.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		liness = append(liness, line+"\n")
	}

	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	for _, i := range liness {
		fmt.Println(i)
	}
}
