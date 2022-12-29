package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("./day-2/input-day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var score = 0
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		var result = int(play(text)) + matchScore(string(text[2]))
		fmt.Printf("Score for input: %s is: %d", text, result)
		score += result
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
	}
	readFile.Close()

	fmt.Printf("The Answer is %d \n", score)
}

type Outcome int64

const (
	rock     Outcome = 1
	paper    Outcome = 2
	scissors Outcome = 3
)

func matchScore(input string) int {
	switch input {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		return 0
	}
}

func play(input string) Outcome {
	switch input {
	case "A X":
		return scissors
	case "A Y":
		return rock
	case "A Z":
		return paper
	case "B X":
		return rock
	case "B Y":
		return paper
	case "B Z":
		return scissors
	case "C X":
		return paper
	case "C Y":
		return scissors
	case "C Z":
		return rock
	default:
		fmt.Printf("Error, outside switch: %s", input)
		return 0
	}
}
