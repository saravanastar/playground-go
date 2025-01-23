package calculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	result := calculate("1 + 2")
	if result != 3 {
		t.Errorf("Invalid output %v", result)
	}
}

func TestCalculate2(t *testing.T) {
	result := calculate(" 2-1 + 2 ")
	if result != 3 {
		t.Errorf("Invalid output %v", result)
	}
}

func TestCalculate3(t *testing.T) {
	result := calculate("(1+(4+5+2)-3)+(6+8)")
	if result != 23 {
		t.Errorf("Invalid output %v", result)
	}
}

func TestCalculate4(t *testing.T) {
	result := calculate("2147483647")
	if result != 2147483647 {
		t.Errorf("Invalid output %v", result)
	}
}

func TestCalculate5(t *testing.T) {
	result := calculate("1-(     -2)")
	if result != 3 {
		t.Errorf("Invalid output %v", result)
	}
}
