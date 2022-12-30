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

	var first string
	var second string
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		if len(first) == 0 {
			first = line
		} else if len(second) == 0 {
			second = line
		} else {
			var third = line
			var found = false
			for _, x := range first {
				if found {
					break
				}
				for _, y := range second {
					if found {
						break
					}
					if x == y {
						for _, z := range third {
							if found {
								break
							}
							if x == z {
								items = append(items, x)
								score += alphabetIndex(x)
								fmt.Printf("Found %c as badge, with score %d  \n", x, alphabetIndex(x))
								found = true
							}
						}
					}
				}
			}
			first = ""
			second = ""
		}
	}
	readFile.Close()

	fmt.Printf("The Answer is %d \n", score)
}

func alphabetIndex(i rune) int {
	var foo = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(foo, string(i)) + 1
}
