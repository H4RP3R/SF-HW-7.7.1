package calc

import (
	"fmt"
	"testing"
)

func Test_calculator_Calculate(t *testing.T) {
	tests := []struct {
		a, b    float64
		op      string
		want    float64
		wantErr error
	}{
		{3, 2, "+", 5, nil},
		{-9, -2, "+", -11, nil},
		{-2, 6, "*", -12, nil},
		{1.23, 6.5, "*", 7.995, nil},
		{5, 5, "+", 10, nil},
		{-8, 3, "+", -5, nil},
		{4, 9, "*", 36, nil},
		{-2.5, 5, "*", -12.5, nil},
		{8, 0, "/", 0, ErrDivByZero},
		{10, 2, "/", 5, nil},
		{15, 3, "/", 5, nil},
		{7.5, 1.5, "/", 5, nil},
		{666, 13, "+-", 0, ErrUnknownOp},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v %s %v", tt.a, tt.op, tt.b)
		t.Run(name, func(t *testing.T) {
			c := calculator{}
			got, err := c.Calculate(tt.a, tt.b, tt.op)
			if err != tt.wantErr {
				t.Errorf("calculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
