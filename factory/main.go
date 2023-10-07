package main

import (
	"SDP-assignment-2/factory/computer"
	"fmt"
)

func main() {
	desktop := computer.GetComputer(computer.DesktopType, "8GB", "1TB")
	laptop := computer.GetComputer(computer.LaptopType, "16GB", "512GB SSD")
	server := computer.GetComputer(computer.ServerType, "64GB", "10TB HDD")

	printDetails(desktop)
	printDetails(laptop)
	printDetails(server)
}

func printDetails(c computer.Computer) {
	fmt.Printf("Type: %s, RAM: %s, Storage: %s\n", c.GetDeviceType(), c.GetRAM(), c.GetStorage())
}
