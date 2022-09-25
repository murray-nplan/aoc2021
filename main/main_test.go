package main

import (
	"murray.nplan/day1"
	"testing"
)

func TestDay1(t *testing.T) {
	v1, v2 := day1.Run("../day1/day1.test")
	if v1 != 7 || v2 != 5 {
		t.Fatalf("Wrong values!")
	}
}
