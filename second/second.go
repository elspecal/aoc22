package second

import (
	"log"
	"os"
	"strings"
)

var ruleA = map[[2]string]int{
	{"A", "X"}: 4,
	{"A", "Y"}: 8,
	{"A", "Z"}: 3,
	{"B", "X"}: 1,
	{"B", "Y"}: 5,
	{"B", "Z"}: 9,
	{"C", "X"}: 7,
	{"C", "Y"}: 2,
	{"C", "Z"}: 6,
}

var ruleB = map[[2]string]int{
	{"A", "X"}: 3,
	{"A", "Y"}: 4,
	{"A", "Z"}: 8,
	{"B", "X"}: 1,
	{"B", "Y"}: 5,
	{"B", "Z"}: 9,
	{"C", "X"}: 2,
	{"C", "Y"}: 6,
	{"C", "Z"}: 7,
}

func parse(path string) []string {
	rawinput, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(rawinput)), "\n")
}

func getTotalScore(path string, rule map[[2]string]int) int {
	var score int
	strategy := parse(path)

	for _, round := range strategy {
		choices := (*[2]string)(strings.Split(round, " "))
		score += rule[*choices]
	}

	return score
}

func GetTotalScoreA(path string) int {
	return getTotalScore(path, ruleA)
}

func GetTotalScoreB(path string) int {
	return getTotalScore(path, ruleB)
}
