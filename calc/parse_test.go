package calc

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct{
		expr string
		expect float64
	}{
		{
			"1+3",
			4,
		},
		{
			"3*(7+8)",
			45,
		},
		{
			"1*((2+3)*4)",
			20,
		},
		{
			"1.12+3.2",
			4.32,
		},
		{
			"2*(1.6+(2.7-0.2))",
			8.2,
		},
	}

	for _, test := range tests {
		r := NewCalcParser(test.expr)
		if r.Num != test.expect{
			t.Errorf("Calc: %s, Result: %f, Expect: %f", test.expr, r.Num, test.expect)
		}
	}
}