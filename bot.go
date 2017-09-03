package pizzabot

import "github.com/andrewcopp/geometry"

// Solver identifies structs with the ability to solve Pizzabot challenge.
type Solver interface {
	Solve(g *Grid) string
}

// Simple solves the Pizzabot challenge by visiting houses in the order they
// are given regardless of efficiency.
type Simple struct {
	current *geometry.Point
	order   []*geometry.Point
}

// NewSimple creates a Simple struct using a Config struct to specify the order
// the houses are to be visited.
func NewSimple(config *Config) *Simple {
	return &Simple{
		current: geometry.NewPoint(0, 0),
		order:   config.Points,
	}
}

// Solve returns a string denoting the plan needed to execute the solution to
// the Pizzabot challenge.
func (s *Simple) Solve(g *Grid) string {
	solution := ""
	for _, point := range s.order {
		for s.current.X < point.X {
			solution = solution + "E"
			s.current.X++
		}

		for s.current.X > point.X {
			solution = solution + "W"
			s.current.X--
		}

		for s.current.Y < point.Y {
			solution = solution + "N"
			s.current.Y++
		}

		for s.current.Y > point.Y {
			solution = solution + "S"
			s.current.Y--
		}

		solution = solution + "D"
	}
	return solution
}

// Smart solves the Pizzabot challenge by using A* search to find the most
// efficient route to deliver all of the pizzas.
type Smart struct {
}

// NewSmart returns a Smart struct that can intelligently solve the Pizzabot
// challenge.
func NewSmart() *Smart {
	return &Smart{}
}

// Solve returns a string denoting the plan needed to execute the solution to
// the Pizzabot challenge.
func (s *Smart) Solve(g *Grid) string {
	return "Smart Solution"
}
