package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/cgeorgiades27/aoc2024-go/utils"
)

var re2 = regexp.MustCompile(`\d+`)

const Increasing = 0
const Decreasing = 1

func day2a(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	sum := 0

	for scanner.Scan() {
		safe := true
		line := scanner.Text()
		matches := re2.FindAllString(line, -1)
		direction := Increasing
		for i, match := range matches {
			if i == 0 && len(matches) > 1 {
				diff := utils.Atoi(matches[i+1]) - utils.Atoi(match)
				direction = Increasing
				if diff < 0 {
					direction = Decreasing
				}
				continue
			}

			diff := utils.Atoi(match) - utils.Atoi(matches[i-1])
			tempDirection := Increasing
			if diff < 0 {
				tempDirection = Decreasing
			}
			if tempDirection != direction || utils.Absv(diff) < 1 || utils.Absv(diff) > 3 {
				safe = false
			}
		}
		if safe {
			sum++
		}
	}

	return sum
}

func day2b(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	sum := 0
	badActor := 0

	for scanner.Scan() {
		safe := true
		line := scanner.Text()
		matches := re2.FindAllString(line, -1)
		safe, badActor = checkMatches(matches)

		if safe {
			sum++
		} else {
			matches = append(matches[:badActor], matches[badActor+1:]...)
			safe, badActor = checkMatches(matches)
			if safe {
				sum++
			}
		}
	}

	return sum
}

func checkMatches(matches []string) (bool, int) {
	direction := Increasing
	for i, match := range matches {
		if i == 0 && len(matches) > 1 {
			diff := utils.Atoi(matches[i+1]) - utils.Atoi(match)
			direction = Increasing
			if diff < 0 {
				direction = Decreasing
			}
			continue
		}

		diff := utils.Atoi(match) - utils.Atoi(matches[i-1])
		tempDirection := Increasing
		if diff < 0 {
			tempDirection = Decreasing
		}

		if tempDirection != direction || utils.Absv(diff) < 1 || utils.Absv(diff) > 3 {
			if tempDirection != direction {
				return false, i - 1
			}
			return false, i
		}
	}
	return true, -1
}

func TestDay2a(t *testing.T) {

	infile, err := os.Open("indata/day2")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`),
			2,
		},
		{
			infile,
			257,
		},
	}

	for i, test := range tests {
		actual := day2a(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}

func TestDay2b(t *testing.T) {

	infile, err := os.Open("indata/day2")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`),
			4,
		},
		{
			infile,
			304,
		},
	}

	for i, test := range tests {
		actual := day2b(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}
