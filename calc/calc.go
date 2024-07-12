package calc

import "errors"

var (
	ErrDivByZero = errors.New("division by zero")
	ErrUnknownOp = errors.New("unknown operator")
)

func NewCalculator() calculator {
	return calculator{}
}

// calculator is unnecessary, but it was said so in the task
type calculator struct{}

// Calculate performs the arithmetic operation based on the provided operator.
func (c calculator) Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return c.add(a, b), nil
	case "-":
		return c.sub(a, b), nil
	case "*":
		return c.mul(a, b), nil
	case "/":
		return c.div(a, b)
	}

	return 0, ErrUnknownOp
}

func (c calculator) add(a, b float64) float64 {
	return a + b
}

func (c calculator) sub(a, b float64) float64 {
	return a - b
}

func (c calculator) mul(a, b float64) float64 {
	return a * b
}

func (c calculator) div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}
