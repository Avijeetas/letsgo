package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	Seconds := time.Now().Unix()

	rand.Seed(Seconds)

	target := rand.Intn(100) + 1 // between to 0 to 99

	fmt.Println("I have chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)

	if target < guess {
		fmt.Println("Guess is higher than target")
	} else if target > guess {
		fmt.Println("Guess is lower than target")
	} else {
		fmt.Println("Guess is matched with target")
	}
}
