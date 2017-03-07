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

}

func TestCanMultiple2Numbers(t *testing.T) {

}

func TestCanDivide2Numbers(t *testing.T) {

}
