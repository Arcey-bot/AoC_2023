package ezfile

import (
	"bufio"
	"os"
	"strings"
)

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
