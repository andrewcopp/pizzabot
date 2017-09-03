package pizzabot

import (
	"testing"

	"github.com/andrewcopp/geometry"
)

func TestSimpleSolve(t *testing.T) {
	size := geometry.NewSize(2, 2)
	pointA := geometry.NewPoint(1, 1)
	pointB := geometry.NewPoint(0, 0)
	cases := []struct {
		in  *Config
		out string
	}{
		{&Config{Size: size, Points: []*geometry.Point{pointA}}, "END"},
		{&Config{Size: size, Points: []*geometry.Point{pointB}}, "D"},
		{&Config{Size: size, Points: []*geometry.Point{pointA, pointB}}, "ENDWSD"},
	}

	for _, c := range cases {
		bot := NewSimple(c.in)
		got, _ := bot.Solve(NewGrid(c.in))
		if got != c.out {
			t.Error("Unexpected output.")
		}
	}
}

func TestSmartSolve(t *testing.T) {
	size := geometry.NewSize(2, 2)
	pointA := geometry.NewPoint(1, 1)
	pointB := geometry.NewPoint(0, 0)
	points := []*geometry.Point{pointA, pointB}
	config := &Config{Size: size, Points: points}
	bot := NewSmart()
	got, _ := bot.Solve(NewGrid(config))
	if got != "Smart Solution" {
		t.Error("Unexpected output.")
	}
}
