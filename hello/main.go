package main

import "fmt"
import "hello/greetings"
import "calcul"
import "log"

func main() {
	greet, errGreet := greetings.Greetings("Toto")
	if errGreet != nil { log.Fatal(errGreet) }

	op, errOp := calcul.Operations( 2, 2, "/" )
	if errOp != nil { log.Fatal(errOp) }

	fmt.Println( greet )
	fmt.Println( op )
}
