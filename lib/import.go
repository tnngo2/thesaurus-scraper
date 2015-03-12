package scrape

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ImportWordList(input string) []string {

	var slice []string
	file, err := os.Open(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		var line string
		line = scanner.Text()
		if strings.TrimSpace(line) != "" {
			slice = append(slice, strings.TrimSpace(line))
		}
	}

	return slice
}

func PrintSlice(slice []string) {
	for _, elem := range slice {
		fmt.Println(elem)
	}
}

func PrintSliceHtml(slice []string) string {
	var result string
	for _, elem := range slice {
		result += elem + "<br>"
	}

	return result
}
