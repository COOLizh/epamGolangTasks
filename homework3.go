package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

/*Task 1 : average([...]int) float64 function that returns
an average value of array (sum / N)*/
func average(array []int) float64 {
	if len(array) == 0 {
		return 0
	}
	var sum int
	for _, item := range array {
		sum += item
	}
	return float64(sum) / float64(len(array))
}

/*Task 2.1 : function that returns the longest word from the slice
of strings (the first if there are more than one).*/
func max(array []string) string {
	if len(array) == 0 {
		return "There is no strings in slice!"
	}
	maxString := array[0]
	for i := 1; i < len(array); i++ {
		if utf8.RuneCountInString(array[i]) > utf8.RuneCountInString(maxString) {
			maxString = array[i]
		}
	}
	return maxString
}

/*Task 2.2 : function that returns the copy of the original
slice in reverse order. The type of elements is int64.*/
func reverse(array []int64) []int64 {
	slice := make([]int64, 0, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		slice = append(slice, array[i])
	}
	return slice
}

/*Task 3 : function that prints map values sorted in order of increasing keys.*/
func printSorted(m map[int]string) {
	keys := make([]int, 0, len(m))
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, item := range keys {
		fmt.Print(m[item], " ")
	}
	fmt.Println()
}

func main() {
	/*Task 1 : Implement average([...]int) float64 function
	that returns an average value of array (sum / N)*/
	fmt.Println(average([]int{1, 2, 3, 4, 5, 6}))
	var firstTaskSlice []int
	fmt.Println(average(firstTaskSlice))
	fmt.Println(average([]int{-1, 3, -2, -4}))

	/*Task 2.1 : Write max([]string) string function that returns
	the longest word from the slice of strings (the first if there are more than one).*/
	fmt.Println(max([]string{"one", "two", "three"}))
	fmt.Println(max([]string{"one", "two"}))
	var firstSlice []string
	fmt.Println(max(firstSlice))

	/*Task 2.2 : Write reverse([]int64) []int64 function that returns
	the copy of the original slice in reverse order. The type of elements is int64.*/
	var secondSlice = []int64{1, 2, 5, 15}
	fmt.Println(reverse(secondSlice))
	fmt.Println(secondSlice)

	/*Task 3 : Implement printSorted(map[int]string)function
	that prints map values sorted in order of increasing keys.*/
	var firstM = map[int]string{2: "a", 0: "b", 1: "c"}
	var secondM = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(firstM)
	printSorted(secondM)
}
