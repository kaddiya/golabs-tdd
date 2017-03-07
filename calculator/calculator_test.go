package calculator

import "testing"

func TestCanAdd2Numbers(t *testing.T) {
	result := Add(5, 6)
	logAndAssert(11, result, "Addition of 5 and 6 should yield 11 but yielded", t)
}

func TestCanSubtract2Numbers(t *testing.T) {
	result := Subtract(5, 4)
	logAndAssert(1, result, "subtraction of 4 from 5 should yield 1 but yielded ", t)
}

func TestCanMultiple2Numbers(t *testing.T) {
	result := Multiply(10, 9)
	logAndAssert(90, result, "Multiplication of 10 and 9 should yield 90 but yielded ", t)
}

func TestCanDivide2Numbers(t *testing.T) {

}

func logAndAssert(expected, actual int, message string, t *testing.T) {
	if actual != expected {
		t.Log(message, actual)
		t.Fail()
	}
}
