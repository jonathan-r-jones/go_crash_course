package main

import "testing"

func Test_add(t *testing.T) {
	a, b, c := 2, 2, 4

	result := add(a, b)
	if result != c {
		t.Fail()
	} else {
		t.Logf("success")
	}

}