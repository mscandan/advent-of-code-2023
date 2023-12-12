package day04

import (
	"fmt"
	"math"
	"strings"

	"github.com/mscandan/advent-of-code-2023/utils"
	"golang.org/x/exp/slices"
)

func part1() {
	total := 0
	lines := utils.ReadFile("day04/input.txt")

	for _, line := range lines {
		game_splitted := strings.Split(line, ":")
		cards := strings.Split(strings.Trim(game_splitted[1], " "), "|")

		winning_cards := strings.Split(strings.Trim(cards[0], " "), " ")
		own_cards := strings.Split(strings.Trim(cards[1], " "), " ")

		total_matches := 0
		unique_own_cards := make([]string, 0)
		unique_winning_cards := make([]string, 0)

		for _, card := range winning_cards {
			if card != "" && !slices.Contains(unique_winning_cards, card) {
				unique_winning_cards = append(unique_winning_cards, card)
			}
		}

		for _, card := range own_cards {
			if card != "" && !slices.Contains(unique_own_cards, card) {
				unique_own_cards = append(unique_own_cards, card)
			}
		}

		for _, card := range unique_own_cards {
			if slices.Contains(unique_winning_cards, card) {
				total_matches++
			}
		}

		if total_matches > 0 {
			total += int(math.Pow(2, float64(total_matches-1)))
		}

	}

	fmt.Printf("Solution for Day 4 Part 1: %d\n", total)
}

func part2() {
	total := 0
	cards_counts := map[int]int{}

	lines := utils.ReadFile("day04/input.txt")

	for idx, line := range lines {
		curr, ok := cards_counts[idx]
		if !ok {
			cards_counts[idx] = 1
		} else {
			cards_counts[idx] = curr + 1
		}

		game_splitted := strings.Split(line, ":")
		cards := strings.Split(strings.Trim(game_splitted[1], " "), "|")

		winning_cards := strings.Split(strings.Trim(cards[0], " "), " ")
		own_cards := strings.Split(strings.Trim(cards[1], " "), " ")

		total_matches := 0
		unique_own_cards := make([]string, 0)
		unique_winning_cards := make([]string, 0)

		for _, card := range winning_cards {
			if card != "" && !slices.Contains(unique_winning_cards, card) {
				unique_winning_cards = append(unique_winning_cards, card)
			}
		}

		for _, card := range own_cards {
			if card != "" && !slices.Contains(unique_own_cards, card) {
				unique_own_cards = append(unique_own_cards, card)
			}
		}

		for _, card := range unique_own_cards {
			if slices.Contains(unique_winning_cards, card) {
				total_matches++
			}
		}

		if total_matches > 0 {
			curr, _ := cards_counts[idx]
			for i := 0; i < total_matches; i++ {
				temp, ok := cards_counts[idx+1+i]
				if !ok {
					cards_counts[idx+1+i] = curr
				} else {
					cards_counts[idx+1+i] = temp + curr
				}
			}
		}
	}

	for _, value := range cards_counts {
		total += value
	}

	fmt.Printf("Solution for Day 4 Part 2: %d\n", total)
}

func Solution() {
	part1()
	part2()
}
