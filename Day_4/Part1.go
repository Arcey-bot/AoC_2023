package main	

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
)

type Card struct {
	id int
	winners, selection []int
}

const FILE string = "input.txt"
func main() {
	data := to_cards(load_data(FILE))
	for _, v := range data {
		fmt.Println(v)
	}
}

func card_value(card Card) (val int) {


	// value is defined as: 2^0 = 1, 2^1 = 2, 2^2= 4, ...
	return int(math.Pow(2, float64(val)))
}

func to_cards(in [][]string) []Card {
	out := make([]Card, len(in))

	for i, r := range in {
		c := Card{
			i+1,
			to_ints(strings.Split(strings.Trim(r[0], " "), " ")),
			to_ints(strings.Split(strings.Trim(r[1], " "), " ")),
		}
		out = append(out, c)
	}

	return out
}

func to_ints(in []string) []int {
	out := make([]int, len(in))

	for i, n := range in {
		out[i], _ = strconv.Atoi(n)
	}

	return out
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
