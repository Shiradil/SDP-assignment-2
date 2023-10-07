package main

import (
	"SDP-assignment-2/decorator/beverage"
	"SDP-assignment-2/decorator/coffee"
	"SDP-assignment-2/decorator/decorators"
	"fmt"
)

func main() {
	basicCoffee := coffee.NewBasicCoffee()

	var milkCoffee beverage.Beverage = decorators.NewMilkDecorator(basicCoffee)
	var sugarMilkCoffee beverage.Beverage = decorators.NewSugarDecorator(milkCoffee)

	fmt.Println(basicCoffee.Description(), ":", basicCoffee.Cost())
	fmt.Println(milkCoffee.Description(), ":", milkCoffee.Cost())
	fmt.Println(sugarMilkCoffee.Description(), ":", sugarMilkCoffee.Cost())
}
