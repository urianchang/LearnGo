package main

import (
	"fmt"
)

type Duck struct {
	ID 		int64
	Name	string
}

func (d *Duck) Eat() {
	fmt.Printf("Duck %s eats!\n", d.Name)
}

type Mallard struct {
	Duck
	Color 	string
}

func (m *Mallard) Sleep() {
	fmt.Printf("Mallard %s sleeps!\n", m.Name)
}

func main() {
	duck := new(Duck)
	duck.ID = 1
	duck.Name = "Pickles"

	mallard := new(Mallard)
	mallard.Color = "Green"
	mallard.Duck = *duck

	duck.Eat()		// Duck Pickles eats!
	mallard.Eat()	// Duck Pickles eats!
	mallard.Sleep()	// Mallard Pickles eats!
}