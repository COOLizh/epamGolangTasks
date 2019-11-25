package tests

import (
	"testing"
	"time"

	"github.com/COOLizh/epam/epamGolangTasks/hw4/tasks"
	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {
	assert := assert.New(t)
	ivanIvanovDate, _ := time.Parse("2006-Jan-02", "2005-Aug-10")
	artiomIvanovDate, _ := time.Parse("2006-Jan-02", "2005-Aug-10")
	ivanIvanovDate2, _ := time.Parse("2006-Jan-02", "2003-Aug-05")
	p := tasks.People{
		{FirstName: "I", LastName: "Ivanov", BirthDay: ivanIvanovDate},
		{FirstName: "a", LastName: "Ivanov", BirthDay: artiomIvanovDate},
		{FirstName: "I", LastName: "Ivanov", BirthDay: ivanIvanovDate2},
	}
	assert.Equal(p.Less(0, 1), true, "ERROR: check 0 and 1 element of People")
	assert.Equal(p.Less(0, 0), false, "ERROR: check the same People")
	assert.Equal(p.Less(0, 2), true, "ERROR: check 0 and 2 element of People")
	assert.Equal(p.Less(1, 2), true, "ERROR: check 1 and 2 element of People")
}
