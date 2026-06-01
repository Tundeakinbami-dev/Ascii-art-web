package main

import (
	"fmt"
	"os"
	"strings"
)
func loadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dataLines := strings.Split(string(data), "\n")

	if len(dataLines) < 800 {
		fmt.Println("banner file incomplete")
	}

	Ascii := make(map[rune][]string)

	index := 1

	for i := 32; i <= 126; i++ {
		Ascii[rune(i)] = dataLines[index:index+8]
		index += 9
	}
	return Ascii, nil
}