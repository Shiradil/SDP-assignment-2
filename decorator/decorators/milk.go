package decorators

import "SDP-assignment-2/decorator/beverage"

type MilkDecorator struct {
	beverage beverage.Beverage
}

func NewMilkDecorator(beverage beverage.Beverage) *MilkDecorator {
	return &MilkDecorator{beverage: beverage}
}

func (m *MilkDecorator) Cost() float64 {
	return 0.50 + m.beverage.Cost()
}

func (m *MilkDecorator) Description() string {
	return m.beverage.Description() + ", Milk"
}
