package main

import "fmt"

// ===================================

type Animal interface {
	Parler()
}

// func FaireParler(a Animal) {
// 	a.Parler()
// }

// ===================================

type Chien struct { name string }

func (c Chien) Parler() { fmt.Printf( "%v : Woof woof !\n", c.name ) }

// ===================================

type Chat struct { name string }

func (c Chat) Parler() { fmt.Printf( "%v : Miaow Miaow !\n", c.name ) }

// ===================================

func main() {

	objs := []Animal {
		Chien{"Medor"},
		Chat{"Nemo"},
	}

	for _, obj := range objs {
		obj.Parler()
	}
}
