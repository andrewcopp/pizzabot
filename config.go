package pizzabot

import (
	"errors"
	"strconv"
	"strings"

	"github.com/andrewcopp/geometry"
)

// Parser identies objects that can parse the Pizzabot args.
type Parser interface {
	Parse(args []string) error
}

// Config is a strongly-typed representation of the command line args.
type Config struct {
	Size   *geometry.Size
	Points []*geometry.Point
}

// Parse takes the args and loads them into the Config struct.
func (c *Config) Parse(args []string) error {
	if len(args) != 2 {
		return errors.New("Expected one arg")
	}

	args = strings.Split(args[1], " (")

	var err error

	c.Size, err = c.dimensions(args[0])
	if err != nil {
		return err
	}

	c.Points, err = c.coordinates(args[1:])
	if err != nil {
		return err
	}

	if err := c.check(); err != nil {
		return err
	}

	return nil
}

// Dimensions is a helper function that parses an argument into a Size struct.
func (c *Config) dimensions(arg string) (*geometry.Size, error) {
	values := strings.Split(arg, "x")

	if len(values) != 2 {
		return nil, errors.New("Expected size to be two-dimensional")
	}

	w, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, err
	}

	h, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, err
	}

	return geometry.NewSize(w, h), nil
}

// Coordinates is a helper function that parses args into a slice of Point
// structs.
func (c *Config) coordinates(coords []string) ([]*geometry.Point, error) {

	points := []*geometry.Point{}
	for _, coord := range coords {
		point := geometry.NewPoint(0, 0)
		cmpnts := strings.Split(strings.TrimSpace(coord), ", ")

		var err error
		if len(cmpnts) != 2 {
			text := "Expected coordinates to be in two-dimensional space"
			return nil, errors.New(text)
		}

		point.X, err = strconv.Atoi(cmpnts[0])
		if err != nil {
			return nil, err
		}

		point.Y, err = strconv.Atoi(strings.Trim(cmpnts[1], ")"))
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}

func (c *Config) check() error {

	for _, point := range c.Points {
		if point.X >= c.Size.Width || point.X <= -1 {
			return errors.New("House exists outside of map")
		}

		if point.Y >= c.Size.Height || point.Y <= -1 {
			return errors.New("House exists outside of map")
		}
	}

	return nil
}
