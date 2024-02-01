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

func load_data(f string) (lines [][]string) {
	file, err := os.Open(f)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return
}
