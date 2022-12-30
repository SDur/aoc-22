package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("./day-3/input-day3.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var score = 0
	var items []rune
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var firstHalf = line[:len(line)/2]
		var secondHalf = line[len(line)/2:]
		fmt.Printf("\t first: %s second: %s  \n", firstHalf, secondHalf)

		var found = false
		for _, x := range firstHalf {
			if found {
				break
			}
			for _, y := range secondHalf {
				if found {
					break
				}
				if x == y {
					items = append(items, x)
					score += alphabetIndex(x)
					fmt.Printf("Found %c as double, with score %d  \n", x, alphabetIndex(x))
					found = true
				}
			}
		}

	}
	readFile.Close()

	fmt.Printf("The Answer is %d \n", score)
}

func alphabetIndex(i rune) int {
	var foo = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(foo, string(i)) + 1
}
