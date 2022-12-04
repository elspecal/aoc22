package first

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calcOrderedSumOfCaloriesPerElf(path string) []int {
	rawinput, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	inventory := strings.Split(string(rawinput), "\n\n")
	calorysums := make([]int, len(inventory))

	for i, elf := range inventory {
		var sum int

		for _, entry := range strings.Split(elf, "\n") {
			calory, _ := strconv.Atoi(entry)
			sum += calory
		}

		calorysums[i] = sum
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorysums)))

	return calorysums
}

func GetFirst(path string) int {
	return calcOrderedSumOfCaloriesPerElf(path)[0]
}

func GetSumTopThree(path string) int {
	var result int

	for _, sum := range calcOrderedSumOfCaloriesPerElf(path)[:3] {
		result += sum
	}

	return result
}
