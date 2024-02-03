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
	id, mc, value int
	matches map[int]int
}

const FILE string = "input.txt"
func main() {
	data := to_cards(load_data(FILE))
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println(sum_cards(data))
}

func sum_cards(in []Card) (s int) {
	for _, c := range in {
		s += c.value
	}
	return 
}

func to_cards(in [][]string) []Card {
	out := make([]Card, len(in))

	for i, r := range in {
		winners, selected := strings.Fields(strings.Trim(r[0], " ")), strings.Fields(strings.Trim(r[1], " "))
		win := make(map[int]int, len(winners))
		matches := 0
		// fmt.Printf("Converting W-%q, S-%q, W-%v\n",winners, selected, win )
		
		for _, w := range winners { // save numbers we're looking for
			n, _ := strconv.Atoi(w)
			win[n] = 0
		}

		for _, v := range selected { // save selected values and number that won
			n, _ := strconv.Atoi(v)
			_, exists := win[n]
			if exists {
				matches++
				win[n]++
			}
		}

		card_value := 0
		if matches >= 0 { // value is only nonzero when there are matches
			// 2^m defines card value, 2^0 is the value of 1 number matching, 2^1 is 2 matches, etc.
			card_value = int(math.Pow(2, float64(matches-1)))
		}

		// fmt.Printf("Converted M-%v, W-%v\n", matches, win)
		out[i] = Card{
			id: i+1,
			mc: matches,
			value: card_value,
			matches: win,
		}
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
