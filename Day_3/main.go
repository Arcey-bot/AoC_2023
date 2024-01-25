package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILE string = "input1.txt"
const SKIP string = "."

type Schematic struct {
	scheme [][]string
}

func (s *Schematic) String() string {
	str := strings.Builder{}

	for _, c := range s.scheme {
		for _, r := range c {
			str.WriteString(r)
		}
		str.WriteRune('\n')
	}

	return str.String()
}

// possible part(s) contained in the adjacent tiles
func (s *Schematic) part_square(r, c int) bool {
	digit, _ := is_digit(s.scheme[r][c])
	return s.scheme[r][c] != SKIP && !digit
}

// consume digits horizontally in the given position to determine the part's value
func (s *Schematic) part_value(r, c int) int {
	numeric, num := is_digit(s.scheme[r][c])

	if !numeric || num == nil {
		return 0 // err
	}

	base := 10
	n := 0

	// go from c to end of array
	for i := c; i < len(s.scheme[r]); i++ {
		numeric, num = is_digit(s.scheme[r][i])
		// fmt.Printf("Received Digit: %v - %v\n", numeric, *num)

		if numeric {
			fmt.Printf("[F] Number: (%v * %v) + %v = %v\n", n, base, *num, (n*base)+*num)
			s.scheme[r][i] = SKIP // consume
			n = (n * base) + *num // mult by 10 + new val is equivalent to appending a number
		} else {
			break // no more numbers
		}
	}

	fmt.Printf("Build now at %v\n", n)
	base = 1 // to build from back to front, the base must begin at 1 and grow by 10 per digit

	// go from c-1 to beginning
	for i := c - 1; i >= 0; i-- {
		numeric, num = is_digit(s.scheme[r][i])

		if numeric {
			fmt.Printf("[B] Number: (%v * %v) + %v = %v\n", *num, base, n, (*num*base)+n)
			s.scheme[r][i] = SKIP // consume
			n += *num * base      // insert number at beginning
			base *= 10
		} else {
			break // no more numbers
		}
	}

	return n
}

// parse and sum any numbers in perimeter
func (s *Schematic) parse_perimeter(r, c int) (sum int) {
	// TL
	if s.inbounds(r-1, c-1) {
		sum += s.part_value(r-1, c-1)
		fmt.Printf("Value @ (%v, %v) %v\n", r-1, c-1, sum)
	}
	// Middle
	if s.inbounds(r-1, c) {
		sum += s.part_value(r-1, c)
		fmt.Printf("Value @ (%v, %v) %v\n", r-1, c, sum)
	}
	// TR
	if s.inbounds(r-1, c+1) {
		sum += s.part_value(r-1, c+1)
		fmt.Printf("Value @ (%v, %v) %v\n", r-1, c+1, sum)
	}
	// L
	if s.inbounds(r, c-1) {
		sum += s.part_value(r, c-1)
		fmt.Printf("Value @ (%v, %v) %v\n", r, c-1, sum)
	}
	// R
	if s.inbounds(r, c+1) {
		sum += s.part_value(r, c+1)
		fmt.Printf("Value @ (%v, %v) %v\n", r, c+1, sum)
	}
	// BL
	if s.inbounds(r+1, c-1) {
		sum += s.part_value(r+1, c-1)
		fmt.Printf("Value @ (%v, %v) %v\n", r+1, c-1, sum)
	}
	// BM
	if s.inbounds(r+1, c) {
		sum += s.part_value(r+1, c)
		fmt.Printf("Value @ (%v, %v) %v\n", r+1, c, sum)
	}
	// BR
	if s.inbounds(r+1, c+1) {
		sum += s.part_value(r+1, c+1)
		fmt.Printf("Value @ (%v, %v) %v\n", r+1, c+1, sum)
	}
	return sum
}

func (s *Schematic) inbounds(r, c int) bool {
	return (0 <= r && r <= len(s.scheme) &&
		0 <= c && c <= len(s.scheme[0]))
}

func (s *Schematic) solve() (sum int) {
	for i, r := range s.scheme {
		for j := range r {
			if s.part_square(i, j) {
				// fmt.Printf("%v @ (%v,÷ %v) is a potential part number\n", c, i, j)
				sum += s.parse_perimeter(i, j)
				fmt.Printf("Sum total: %v\n", sum)
			}
		}
	}

	return
}

func main() {
	scheme := Schematic{load_data(FILE)}

	fmt.Println(&scheme)
	fmt.Println(scheme.solve())
	fmt.Println(&scheme)
}

func is_digit(s string) (bool, *int) {
	n, err := strconv.Atoi(s)
	if err == nil { // is a number
		return true, &n
	}
	return false, nil
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
