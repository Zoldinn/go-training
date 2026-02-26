package main

import "fmt"

// ===================================

type Animal interface {
	Parler() string
}

func FaireParler(a Animal) string {
	return a.Parler()
}

// ===================================

type Chien struct {
	name string
}

func (c Chien) Parler() string {
	return "Woof woof !"
}

// ===================================

func (c Chat) Parler() string {
	return "Miaow Miaow !"
}

type Chat struct {
	name string
}

// ===================================

func main() {
	var a Animal
	chien := Chien{"boby"}
	chat := Chat{"toto"}

	a = chien
	fmt.Println( FaireParler(a) )
	a = chat
	fmt.Println( FaireParler(a) )
}
