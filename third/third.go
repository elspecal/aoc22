package third

import (
	"log"
	"os"
	"strings"
)

var prioritymap = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9,
	"j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17,
	"r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25,
	"z": 26, "A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33,
	"H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41,
	"P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49,
	"X": 50, "Y": 51, "Z": 52,
}

func parse(path string) []string {
	rawinput, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(rawinput)), "\n")
}

func findPriority(rucksack string) int {
	var priority int

	items := strings.Split(rucksack, "")
	capacity := len(rucksack)
	head := items[:capacity/2]
	tail := strings.Join(items[capacity/2:], "")

	for _, item := range head {
		if strings.Contains(tail, item) {
			priority = prioritymap[item]
			break
		}
	}

	return priority
}

func SumPriorities(path string) int {
	var sum int

	for _, rucksack := range parse(path) {
		sum += findPriority(rucksack)
	}

	return sum
}
