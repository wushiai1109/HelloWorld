package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	//got := 15
	//got := 5
	want := 15

	t.Log(got)
	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
