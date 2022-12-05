package fourth

import (
	"regexp"
	"strconv"

	"github.com/elspecal/aoc22/common"
)

func contained(pairs string) bool {
	sep := regexp.MustCompile("[,-]")
	parsed := sep.Split(pairs, -1)

	var sections [4]int
	for i, s := range parsed {
		sections[i], _ = strconv.Atoi(s)
	}

	return (sections[0] >= sections[2] && sections[1] <= sections[3]) ||
		(sections[0] <= sections[2] && sections[1] >= sections[3])
}

func CountContaines(path string) int {
	var count int

	assignments := common.Parse(path)

	for _, pairs := range assignments {
		if contained(pairs) {
			count++
		}
	}

	return count
}
