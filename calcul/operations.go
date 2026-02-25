package calcul

import "errors"

func Operations( a int, b int, op string ) (int, error) {
	switch op {
		case "+" : return a + b, nil
		case "-" : return a - b, nil
		case "*" : return a * b, nil
		case "/" : return a / b, nil
		case "%" : return a % b, nil
		default : return 0, errors.New("Unknown op")
	}
}