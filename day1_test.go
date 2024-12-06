package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/cgeorgiades27/aoc2024-go/utils"
)

var re1 = regexp.MustCompile(`\d+`)

func day1a(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var list1, list2 []int
	for scanner.Scan() {
		line := scanner.Text()
		matches := re1.FindAllString(line, -1)
		list1 = append(list1, utils.Atoi(matches[0]))
		list2 = append(list2, utils.Atoi(matches[1]))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := range list1 {
		sum += utils.Absv(list2[i] - list1[i])
	}

	return sum
}

func day1b(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	frequencyMap := make(map[int]int)
	var list1 []int

	for scanner.Scan() {
		line := scanner.Text()
		matches := re1.FindAllString(line, -1)
		frequencyMap[utils.Atoi(matches[1])]++
		list1 = append(list1, utils.Atoi(matches[0]))
	}

	sum := 0
	for _, num := range list1 {
		if val, exists := frequencyMap[num]; exists {
			sum += num * val
		}
	}

	return sum
}

func TestDay1a(t *testing.T) {

	infile, err := os.Open("indata/day1")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`),
			11,
		},
		{
			infile,
			2970687,
		},
	}

	for i, test := range tests {
		actual := day1a(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}

func TestDay1b(t *testing.T) {

	infile, err := os.Open("indata/day1")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`),
			31,
		},
		{
			infile,
			23963899,
		},
	}

	for i, test := range tests {
		actual := day1b(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}
