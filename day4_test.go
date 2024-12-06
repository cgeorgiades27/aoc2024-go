package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func day4a(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	var matrix [][]byte

	for scanner.Scan() {
		matrix = append(matrix, scanner.Bytes())
	}

	checkWest := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if j-k < 0 || mtx[i][j-k] != str[k] {
				return false
			}
		}
		return true
	}

	checkEast := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if j+k >= len(mtx[i]) || mtx[i][j+k] != str[k] {
				return false
			}
		}
		return true
	}

	checkNorth := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i-k < 0 || mtx[i-k][j] != str[k] {
				return false
			}
		}
		return true
	}

	checkSouth := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i+k >= len(mtx) || mtx[i+k][j] != str[k] {
				return false
			}
		}
		return true
	}

	checkNortheast := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i-k < 0 || j+k >= len(mtx[i]) || mtx[i-k][j+k] != str[k] {
				return false
			}
		}
		return true
	}

	checkNorthwest := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i-k < 0 || j-k < 0 || mtx[i-k][j-k] != str[k] {
				return false
			}
		}
		return true
	}

	checkSoutheast := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i+k >= len(mtx) || j+k >= len(mtx[i]) || mtx[i+k][j+k] != str[k] {
				return false
			}
		}
		return true
	}

	checkSouthwest := func(i, j int, str string, mtx [][]byte) bool {
		for k := 0; k < len(str); k++ {
			if i+k >= len(mtx) || j-k < 0 || mtx[i+k][j-k] != str[k] {
				return false
			}
		}
		return true
	}

	directions := []func(int, int, string, [][]byte) bool{
		checkNorth,
		checkNortheast,
		checkEast,
		checkSoutheast,
		checkSouth,
		checkSouthwest,
		checkWest,
		checkNorthwest,
	}

	sum := 0
	for outerIndex := range matrix {
		for innerIndex := range matrix[outerIndex] {
			for _, checkFunc := range directions {
				if checkFunc(outerIndex, innerIndex, "XMAS", matrix) {
					sum++
				}
			}
		}
	}

	return sum
}

func TestDay4a(t *testing.T) {
	infile, err := os.Open("indata/day4")
	if err != nil {
		t.Fatal("Couldn't open file", err)
	}

	tests := []struct {
		reader   io.Reader
		expected int
	}{
		{
			strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`),
			18,
		},
		{
			infile,
			257,
		},
	}

	for i, test := range tests {
		actual := day4a(test.reader)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, actual %d", i, test.expected, actual)
		}
	}
}
