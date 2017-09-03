package pizzabot

// House represents a house in the neighborhood that may have placed an order.
type House struct {
	Orders int
}

// NewHouse creates a House struct that has zero outstanding orders.
func NewHouse() *House {
	return &House{
		Orders: 0,
	}
}
