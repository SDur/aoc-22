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
		var result = int(play(text)) + shapeScore(string(text[2]))
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
	win  Outcome = 6
	draw Outcome = 3
	loss Outcome = 0
)

func shapeScore(input string) int {
	switch input {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0
	}
}

func play(input string) Outcome {
	switch input {
	case "A X":
		return draw
	case "A Y":
		return win
	case "A Z":
		return loss
	case "B X":
		return loss
	case "B Y":
		return draw
	case "B Z":
		return win
	case "C X":
		return win
	case "C Y":
		return loss
	case "C Z":
		return draw
	default:
		fmt.Printf("Error, outside switch: %s", input)
		return 0
	}
}
