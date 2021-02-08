package go_cp_utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// SplitNums splits a string into a slice of integers
func SplitNums(s string) []int64 {
	nums := strings.Split(strings.TrimSpace(s), " ")
	var result []int64
	for _, n := range nums {
		result = append(result, ParseInt(n))
	}
	return result
}

// ParseInt converts a string into an integer
func ParseInt(s string) int64 {
	num, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return num
}

// ReadInput reads the standard input and return it as a string
func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}
