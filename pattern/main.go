package main

import (
	"fmt"
)

func main() {
	fmt.Println("Patterns")

	fmt.Println("\nFacade:")
	FacadeExample()

	fmt.Println("\nBuilder:")
	BuilderExample()

	fmt.Println("\nVisitor:")
	VisitorExample()

	fmt.Println("\nCommand:")
	CommandExample()

	fmt.Println("\nChain of responsibility:")
	ChainExample()

	fmt.Println("\nFactory:")
	FactoryExample()

	fmt.Println("\nStrategy:")
	StrategyExample()

	fmt.Println("\nState:")
	StateExample()
}
