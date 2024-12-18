package utils

import "strconv"

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func Absv(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
