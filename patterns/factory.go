package main

import "fmt"

type Shooter interface {
	Shoot()
}

type gun struct {
	name string
}

func (g *gun) Shoot() {
	fmt.Println("Shooting", g.name)
}

func NewAK47() *gun {
	return &gun{
		name: "AK-47",
	}
}

func NewGlock17() *gun {
	return &gun{
		name: "Glock 17",
	}
}

func FactoryExample() {
	ak := NewAK47()
	glock := NewGlock17()

	ak.Shoot()
	glock.Shoot()
}
