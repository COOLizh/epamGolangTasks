package tasks

import (
	"errors"
	"math"
)

// MyStrToInt2 converts string to int
func MyStrToInt2(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("Empty input")
	}
	isPositive := true
	tenDegree := len(s) - 1
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if r[i] == 43 {
				tenDegree--
				continue
			} else if r[i] == 45 {
				isPositive = false
				tenDegree--
				continue
			}
		}
		if r[i] < 48 || r[i] > 57 {
			return 0, errors.New("Invalid input")
		}
	}
	var result int
	idx := 0
	overflowError := errors.New("Overflow")
	if r[0] == 43 || r[0] == 45 {
		idx++
	}
	if isPositive {
		for ; idx < len(s); idx++ {
			result += int(math.Pow10(tenDegree)) * int(r[idx]-48)
			tenDegree--
			if result < 0 {
				return 0, overflowError
			}
		}
	} else {
		for ; idx < len(s); idx++ {
			result -= int(math.Pow10(tenDegree)) * int(r[idx]-48)
			tenDegree--
			if result > 0 {
				return 0, overflowError
			}
		}
	}
	return result, nil
}
