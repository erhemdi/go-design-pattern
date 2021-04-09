package builder

type computer struct {
	processor string
	ram       string
	harddisk  string
	vga       string
	monitor   string
}

func NewComputer(processor, ram, harddisk, vga, monitor string) computer {
	return computer{
		processor: processor,
		ram:       ram,
		harddisk:  harddisk,
		vga:       vga,
		monitor:   monitor,
	}
}
