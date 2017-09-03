package pizzabot

// Grid represents the map of a neighborhood.
type Grid struct {
	Houses [][]*House
}

// NewGrid returns a Grid struct given a Config struct that holds information
// regarding the dimensions of the neighborhood and which houses placed orders.
func NewGrid(config *Config) *Grid {
	houses := make([][]*House, config.Size.Width)
	for i := 0; i < config.Size.Width; i++ {
		houses[i] = make([]*House, config.Size.Height)
	}

	for _, p := range config.Points {
		houses[p.X][p.Y] = NewHouse()
	}

	return &Grid{
		Houses: houses,
	}
}
