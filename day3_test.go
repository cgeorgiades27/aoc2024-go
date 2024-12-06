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

var (
	re3  = regexp.MustCompile(`mul\(\d+,\d+\)`)
	re3a = regexp.MustCompile(`\d+`)
	re3b = regexp.MustCompile(`(don\'t\(\))|(do\(\))`)
)

func day3a(r io.Reader) int {
	sum := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		matches := re3.FindAllString(scanner.Text(), -1)
		for _, match := range matches {
			nums := re3a.FindAllString(match, -1)
			product := utils.Atoi(nums[0]) * utils.Atoi(nums[1])
			sum += product
		}
	}
	return sum
}

func TestDay3a(t *testing.T) {
	infile, err := os.Open("indata/day3")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`),
			161,
		},
		{
			infile,
			257,
		},
	}

	for i, test := range tests {
		actual := day3a(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}
