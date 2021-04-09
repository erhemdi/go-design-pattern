package builder

type computerBuilder struct {
	processor string
	ram       string
	harddisk  string
	vga       string
	monitor   string
}

func NewComputerBuilder() computerBuilder {
	return computerBuilder{}
}

func (c computerBuilder) Build() computer {
	return computer{
		processor: c.processor,
		ram:       c.ram,
		harddisk:  c.harddisk,
		vga:       c.vga,
		monitor:   c.monitor,
	}
}

func (c computerBuilder) SetProcessor(processor string) computerBuilder {
	c.processor = processor
	return c
}

func (c computerBuilder) SetRam(ram string) computerBuilder {
	c.ram = ram
	return c
}

func (c computerBuilder) SetHarddisk(harddisk string) computerBuilder {
	c.harddisk = harddisk
	return c
}

func (c computerBuilder) SetVga(vga string) computerBuilder {
	c.vga = vga
	return c
}

func (c computerBuilder) SetMonitor(monitor string) computerBuilder {
	c.monitor = monitor
	return c
}
