package tests

import (
	"math"
	"testing"

	"github.com/COOLizh/epam/epamGolangTasks/hw4/tasks"
	"github.com/stretchr/testify/assert"
)

func TestCalculation(t *testing.T) {
	assert := assert.New(t)
	var s tasks.Figure
	var c tasks.Figure
	for i := 1; i <= 10; i++ {
		s = tasks.Square{A: float64(i)}
		c = tasks.Circle{Radius: float64(i)}
		assert.Equal(s.Perimeter(), 4*float64(i), "ERROR: check perimeter of square")
		assert.Equal(s.Area(), float64(i*i), "ERROR: check area of square")
		assert.Equal(c.Perimeter(), 2*math.Pi*float64(i), "ERROR: check perimeter of circle")
		assert.Equal(c.Area(), math.Pi*float64(i*i), "ERROR: check area of circle")
	}
}
