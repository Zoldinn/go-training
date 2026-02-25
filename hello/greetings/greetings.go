package greetings

import "fmt"
import "errors"
import "math/rand"

func Greetings( name string ) (string, error) {
	if name == "" {
		return "", errors.New( "No name given to Greetings" )
	}
	strs := []string {
		fmt.Sprintf( "Hello %v !\n", name ),
		fmt.Sprintf( "How are you %v ?\n", name ),
		fmt.Sprintf( "Nice to see you %v !\n", name ),
	}
	return strs[rand.Intn( len(strs) )], nil
}