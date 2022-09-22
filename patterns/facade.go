package main

import "fmt"

type hdd struct{}

func (h *hdd) loadOS() {
	fmt.Println("Loaded OS")
}

type gpu struct{}

func (g *gpu) connect() {
	fmt.Println("GPU connected to monitor")
}

type psu struct{}

func (p *psu) on() {
	fmt.Println("Power turned on")
}

type ComputerFacade struct {
	hdd hdd
	gpu gpu
	psu psu
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		hdd: hdd{},
		gpu: gpu{},
		psu: psu{},
	}
}

func (cf *ComputerFacade) Start() {
	cf.psu.on()
	cf.gpu.connect()
	cf.hdd.loadOS()
}

func FacadeExample() {
	cf := NewComputerFacade()
	cf.Start()
}
