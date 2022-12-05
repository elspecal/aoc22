package common

import (
	"log"
	"os"
	"strings"
)

func Parse(path string) []string {
	rawinput, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.TrimSpace(string(rawinput)), "\n")
}
