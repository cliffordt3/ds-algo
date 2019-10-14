package twonumsum

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10

	output := TwoNumberSum(array, target)
	expected := []int{-1, 11}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
