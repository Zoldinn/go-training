package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("Empty name")
	}

	var msg string = fmt.Sprintf("Hello %s !", name)
	return msg, nil
}
