package calculator

import (
	"fmt"
	"testing"
)

func TestCanAdd2Numbers(t *testing.T) {
	result := Add(5, 6)
	if result != 11 {
		t.Log("Addition of 5 and 6 should yield 11 but yielded", result)
		t.Fail()
	}
}

func TestCanSubtract2Numbers(t *testing.T) {
	result := Subtract(5, 4)
	logAndAssert(1, result, t)
}

func TestCanMultiple2Numbers(t *testing.T) {
	result := Multiply(9, 10)
	logAndAssert(90, result, t)
}

func TestCanDivide2Numbers(t *testing.T) {

}

func logAndAssert(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		msg := fmt.Sprintf("The result expected by the operation was %d but got %d", expected, actual)
		t.Log(msg)
		t.Fail()
	}
}
