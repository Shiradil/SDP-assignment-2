package decorators

import "SDP-assignment-2/decorator/beverage"

type SugarDecorator struct {
	beverage beverage.Beverage
}

func NewSugarDecorator(beverage beverage.Beverage) *SugarDecorator {
	return &SugarDecorator{beverage: beverage}
}

func (s *SugarDecorator) Cost() float64 {
	return 0.20 + s.beverage.Cost()
}

func (s *SugarDecorator) Description() string {
	return s.beverage.Description() + ", Sugar"
}
