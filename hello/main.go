package main

import "fmt"
import "hello/greetings"
import "calcul"
import "log"

func main() {
	var a string = greetings.Greetings("Léon")
	b, err := calcul.Operations( 2, 2, "/" )
	if err == nil {
		fmt.Printf( "%s, %d\n", a, b )
	} else {
		log.Fatal(err)
	}
}
