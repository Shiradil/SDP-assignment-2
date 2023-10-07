package computer

type Computer interface {
	GetDeviceType() string
	GetRAM() string
	GetStorage() string
}

type Desktop struct {
	RAM     string
	Storage string
}

func (d *Desktop) GetDeviceType() string {
	return "Desktop"
}

func (d *Desktop) GetRAM() string {
	return d.RAM
}

func (d *Desktop) GetStorage() string {
	return d.Storage
}

type Laptop struct {
	RAM     string
	Storage string
}

func (l *Laptop) GetDeviceType() string {
	return "Laptop"
}

func (l *Laptop) GetRAM() string {
	return l.RAM
}

func (l *Laptop) GetStorage() string {
	return l.Storage
}

type Server struct {
	RAM     string
	Storage string
}

func (s *Server) GetDeviceType() string {
	return "Server"
}

func (s *Server) GetRAM() string {
	return s.RAM
}

func (s *Server) GetStorage() string {
	return s.Storage
}
