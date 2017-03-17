package calculator

import (
	"errors"
	"reflect"
	"testing"
)

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

var divideTests = []struct {
	dividend int
	divisor  int
	result   int
	message  string
	err      error
}{
	{90, 9, 10, "Division of 90 by 9 doesnt not yield 10", nil},
	{12, 4, 3, "Division of 12 by 4 doesnt not yield 3", nil},
	{45, 7, 6, "Division of 45 by 7 doesnt not yield a round number 6", nil},
	{0, 12, 0, "Division of 0 by 12 should give 0 and a nil error", nil},
	{45, 0, 0, "Divisor by 0 should not be allowed", errors.New("Divisor cant be equal to 0")},
	{0, 0, 0, "Divisor by 0 should not be allowed", errors.New("Divisor cant be equal to 0")},
}

func TestDivide(t *testing.T) {
	for _, tuple := range divideTests {
		actual, err := Divide(tuple.dividend, tuple.divisor)
		if err != nil {
			if !reflect.DeepEqual(err, tuple.err) {
				t.Fail()
				t.Log("expected error was " + tuple.err.Error() + " but got " + err.Error())
			}
		}
		logAndAssert(tuple.result, actual, tuple.message, t)
	}
}

func logAndAssert(expected, actual int, message string, t *testing.T) {
	if actual != expected {
		t.Log(message, actual)
		t.Fail()
	}
}
