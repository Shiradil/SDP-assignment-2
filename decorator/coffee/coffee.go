package coffee

type BasicCoffee struct{}

func NewBasicCoffee() *BasicCoffee {
	return &BasicCoffee{}
}

func (c *BasicCoffee) Cost() float64 {
	return 1.00 // base price for coffee
}

func (c *BasicCoffee) Description() string {
	return "Basic Coffee"
}
