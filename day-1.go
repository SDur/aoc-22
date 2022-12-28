package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	Number   int
	Calories int
}

func main() {

	readFile, err := os.Open("input-day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var elfs []elf
	var elfNr = 1
	var calories = 0
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		if text != "" {
			cals, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			calories += cals
			continue
		} else {
			elfs = append(elfs, elf{elfNr, calories})
			elfNr++
			calories = 0
		}
	}
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].Calories > elfs[j].Calories
	})

	fmt.Printf("least calories is %d \n", elfs[len(elfs)-1].Calories)
	fmt.Printf("The Answer is %d \n", elfs[0].Calories)

	readFile.Close()
}
