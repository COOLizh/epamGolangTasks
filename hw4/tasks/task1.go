package tasks

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

//Person contains of all needful information about person
type Person struct {
	FirstName string
	LastName  string
	BirthDay  time.Time
}

//People is slice of persons
type People []Person

func (slice People) Len() int {
	return len(slice)
}

func (slice People) Less(i, j int) bool {
	if slice[i].BirthDay.Equal(slice[j].BirthDay) {
		if strings.ToLower(slice[i].FirstName) == strings.ToLower(slice[j].FirstName) {
			return strings.ToLower(slice[i].LastName) < strings.ToLower(slice[j].LastName)
		}
		return strings.ToLower(slice[i].FirstName) < strings.ToLower(slice[j].LastName)
	}
	return slice[i].BirthDay.After(slice[j].BirthDay)
}

func (slice People) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice People) String() string {
	var result string
	for _, item := range slice {
		result += fmt.Sprintf("%s %s %s\n", item.FirstName, item.LastName, item.BirthDay)
	}
	return result
}

// Task1 : Implement sort.Interface interface for People type
func Task1() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	artiomIvanovDate, err2 := time.Parse("2006-Jan-02", "2005-Aug-10")
	ivanIvanovDate2, err3 := time.Parse("2006-Jan-02", "2003-Aug-05")
	if err != nil || err2 != nil || err3 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Artiom", "Ivanov", artiomIvanovDate},
		{"Ivan", "Ivanov", ivanIvanovDate2},
	}
	sort.Sort(p)
	fmt.Println(p)
}
