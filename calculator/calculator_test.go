package calculator

import "testing"

func TestCanAdd2Numbers(t *testing.T) {
	result := Add(5, 6)
	if result != 11 {
		t.Log("Addition of 5 and 6 should yield 11 but yielded", result)
		t.Fail()
	}
}

func TestCanSubtract2Numbers(t *testing.T) {
	result := Subtract(5, 4)

	if result != 1 {
		t.Log("subtraction of 4 from 5 should yield 1 but yielded", result)
		t.Fail()
	}
}

func TestCanMultiple2Numbers(t *testing.T) {

}

func TestCanDivide2Numbers(t *testing.T) {

}
