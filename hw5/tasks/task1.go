package tasks

import "strconv"

// MyStrToInt converts string to int
func MyStrToInt(s string) (int, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	return int(num), err
}
