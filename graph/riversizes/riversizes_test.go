package riversizes

import (
	"reflect"
	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{}
	input := [][]int{
		{0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	expected := []int{1}
	input := [][]int{
		{1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	expected := []int{1, 2, 3}
	input := [][]int{
		{1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase4(t *testing.T) {
	expected := []int{1, 1, 2, 3}
	input := [][]int{
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{0, 0, 1, 0},
		{1, 0, 1, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase5(t *testing.T) {
	expected := []int{1, 2, 2, 2, 5}
	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase6(t *testing.T) {
	expected := []int{1, 1, 2, 2, 5, 21}
	input := [][]int{
		{1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0},
		{1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
		{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase7(t *testing.T) {
	expected := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	input := [][]int{
		{1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase8(t *testing.T) {
	expected := []int{1, 1, 1, 1, 1, 1, 1, 1, 7}
	input := [][]int{
		{1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase9(t *testing.T) {
	expected := []int{}
	input := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase10(t *testing.T) {
	expected := []int{49}
	input := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase11(t *testing.T) {
	expected := []int{3, 5, 6}
	input := [][]int{
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 0, 0, 0, 1, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}

func TestCase12(t *testing.T) {
	expected := []int{1, 1, 2, 6, 10}
	input := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}
