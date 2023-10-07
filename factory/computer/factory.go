package computer

type ComputerType string

const (
	DesktopType ComputerType = "Desktop"
	LaptopType  ComputerType = "Laptop"
	ServerType  ComputerType = "Server"
)

func GetComputer(ct ComputerType, ram, storage string) Computer {
	switch ct {
	case DesktopType:
		return &Desktop{RAM: ram, Storage: storage}
	case LaptopType:
		return &Laptop{RAM: ram, Storage: storage}
	case ServerType:
		return &Server{RAM: ram, Storage: storage}
	default:
		return nil
	}
}
