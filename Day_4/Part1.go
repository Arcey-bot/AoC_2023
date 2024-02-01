package main	

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const FILE string = "input.txt"
func main() {
	data := load_data(FILE)
	fmt.Println(data)
}

// index + 1 is card number
func load_data(f string) (lines [][]string) {
	file, err := os.Open(f)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Get s[1] of split on colon since we don't care for card number
	for scanner.Scan() {
		lines = append(lines, strings.Split(strings.Split(scanner.Text(), ":")[1], "|"))
	}

	return lines
}
