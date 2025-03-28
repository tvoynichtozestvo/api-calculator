package logic

import "errors"

func Calculate(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operation: " + op)
	}
}
