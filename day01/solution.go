package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	total := 0
	f, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, "")

		var firstNum int = -1
		var lastNum int = -1
		for _, value := range splitted {
			converted, err := strconv.Atoi(value)

			if err == nil {
				if firstNum == -1 {
					firstNum = converted
				}
				lastNum = converted
			}
		}

		total += firstNum*10 + lastNum

	}

	fmt.Printf("Solution for Day 1 Part 1: %d\n", total)
}

var numStrs = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func part2() {

	total := 0
	f, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, "")

		var firstNum int = -1
		var lastNum int = -1
		for i, value := range splitted {
			converted, err := strconv.Atoi(value)

			if err == nil {
				if firstNum == -1 {
					firstNum = converted
				}
				lastNum = converted
			} else {
				for numStr, num := range numStrs {
					s := line[i:min(i+len(numStr), len(line))]
					if s == numStr {
						if firstNum == -1 {
							firstNum = num
						}
						lastNum = num
					}
				}
			}
		}

		total += firstNum*10 + lastNum

	}

	fmt.Printf("Solution for Day 1 Part 2: %d\n", total)
}

func Solution() {
	part1()
	part2()
}
