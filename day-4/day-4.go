package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	readFile, err := os.Open("./day-4/input-day4.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var score = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var pair = strings.Split(line, ",")
		var first = strings.Split(pair[0], "-")
		var second = strings.Split(pair[1], "-")

		var firstFrom, _ = strconv.Atoi(first[0])
		var firstTo, _ = strconv.Atoi(first[1])

		var secondFrom, _ = strconv.Atoi(second[0])
		var secondTo, _ = strconv.Atoi(second[1])

		if (firstFrom <= secondFrom && firstTo >= secondTo) || (secondFrom <= firstFrom && secondTo >= firstTo) {
			fmt.Printf("Found an overlapping pair: %s \n", line)
			score++
		}
	}
	readFile.Close()
	fmt.Printf("The Answer is %d \n", score)
}
