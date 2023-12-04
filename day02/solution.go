package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)

	counts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, ":")

		currentId, _ := strconv.Atoi(strings.Split(splitted[0], " ")[1])

		gameRecords := strings.Split(splitted[1], ";")

		isValidGame := true
		for _, record := range gameRecords {
			grabbed := strings.Split(record, ",")
			for _, pair := range grabbed {
				pairSplitted := strings.Split(strings.Trim(pair, " "), " ")
				count, _ := strconv.Atoi(pairSplitted[0])
				key := pairSplitted[1]

				if count > counts[key] {
					isValidGame = false
				}

			}

		}

		if isValidGame {
			total += currentId
		}

	}

	fmt.Printf("Solution for Day 2 Part 1: %d\n", total)
}

func part2() {
	f, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		counts := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		splitted := strings.Split(line, ":")

		gameRecords := strings.Split(splitted[1], ";")

		for _, record := range gameRecords {
			grabbed := strings.Split(record, ",")
			for _, pair := range grabbed {
				pairSplitted := strings.Split(strings.Trim(pair, " "), " ")
				count, _ := strconv.Atoi(pairSplitted[0])
				key := pairSplitted[1]

				if count > counts[key] {
					counts[key] = count
				}
			}

		}

		totalPower := 1

		for _, value := range counts {
			totalPower *= value
		}

		total += totalPower

	}

	fmt.Printf("Solution for Day 2 Part 2: %d\n", total)
}

func Solution() {
	part1()
	part2()
}
