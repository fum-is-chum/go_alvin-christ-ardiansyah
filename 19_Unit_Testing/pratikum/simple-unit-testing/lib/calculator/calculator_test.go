package calculator

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	test1 := Addition(1, 2)
	if test1 != 3 {
		t.Error(fmt.Sprintf("Expected 1 + 3 to equal 3, instead of %d", test1))
	}

	test2 := Addition(-1, -2)
	if test2 != -3 {
		t.Error(fmt.Sprintf("Expected (-1) + (-2) to equal (-3), instead of %d", test2))
	}
}

func TestSubtraction(t *testing.T) {
	test1 := Subtraction(1, 2)
	if test1 != -1 {
		t.Error(fmt.Sprintf("Expected 1 - 2 to equal (-1), instead of %d", test1))
	}

	test2 := Subtraction(5, 2)
	if test2 != 3 {
		t.Error(fmt.Sprintf("Expected 5 - 3 to equal 2, instead of %d", test2))
	}
}

func TestDivision(t *testing.T) {
	test1 := Division(8, 2)
	if test1 != 4 {
		t.Error(fmt.Sprintf("Expected 8 / 2 to equal 4, instead of %d", test1))
	}

	test2 := Division(6, -2)
	if test2 != -3 {
		t.Error(fmt.Sprintf("Expected 1/1 to equal -3, instead of %d", test2))
	}
}

func TestMultiply(t *testing.T) {
	test1 := Multiply(2, 2)
	if test1 != 4 {
		t.Error(fmt.Sprintf("Expected 2x2 to equal 4, instead of %d", test1))
	}

	test2 := Multiply(10, -3)
	if test2 != -30 {
		t.Error(fmt.Sprintf("Expected 10x(-3) to equal -30, instead of %d", test2))
	}
}