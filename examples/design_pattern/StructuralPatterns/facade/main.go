package main

import "fmt"

/*
the Facade Pattern
	> It defines a higher-level interface that makes the subsystem easier to use.

*/

type CPU struct{}

func (*CPU) Freeze() {
	fmt.Println("CPU Freeze")
}

func (*CPU) Jump(position int) {
	fmt.Printf("CPU Jump to %d\n", position)
}

func (*CPU) Execute() {
	fmt.Println("CPU Execute")
}

type Memory struct{}

func (*Memory) Load(position int, data string) {
	fmt.Printf("Memory Load data '%s' to position %d\n", data, position)
}

type HardDrive struct{}

func (*HardDrive) Read(position int, size int) string {
	data := fmt.Sprintf("HardDrive Read data from position %d with size %d", position, size)
	fmt.Println(data)
	return data
}

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	c.cpu.Freeze()
	c.memory.Load(0, "boot_loader")
	c.cpu.Jump(0)
	c.cpu.Execute()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
