package main

import "fmt"

type computer struct {
	cpu string
	gpu string
	hdd string
}

func (c computer) String() string {
	return fmt.Sprintf("CPU: %v\nGPU: %v\nHDD: %v",
		c.cpu, c.gpu, c.hdd,
	)
}

type ComputerBuilder interface {
	SetCPU(string) *ComputerBuilder
	SetGPU(string) *ComputerBuilder
	SetHDD(string) *ComputerBuilder
	Build() computer
}

type computerBuilder struct {
	computer *computer
}

func NewComputerBuilder() *computerBuilder {
	return &computerBuilder{
		computer: &computer{},
	}
}

func (cb *computerBuilder) SetCPU(cpu string) *computerBuilder {
	cb.computer.cpu = cpu
	return cb
}

func (cb *computerBuilder) SetGPU(gpu string) *computerBuilder {
	cb.computer.gpu = gpu
	return cb
}

func (cb *computerBuilder) SetHDD(hdd string) *computerBuilder {
	cb.computer.hdd = hdd
	return cb
}

func (cb *computerBuilder) Build() computer {
	return *cb.computer
}

func BuilderExample() {
	builder := NewComputerBuilder()

	builder.SetCPU("Intel Core I7 7700K")
	builder.SetGPU("NVIDIA GTX 1080")
	builder.SetHDD("1024 GB")

	fmt.Println(builder.Build())
}
