package main

import "testing"

type test struct {
	data []int
	answer int
}


func TestAllTesteSoma(t *testing.T) {
	tests := []test {
		test{data: []int{50,50}, answer: 100},
		test{data: []int{30,50}, answer: 80},
		test{data: []int{12,80}, answer: 93},
		test{data: []int{32,32}, answer: 64},
	}

	for _, v := range tests {
		z := sum(v.data...)

		if z != v.answer {
			t.Error("Expected: ", v.answer, "but got: ", z)	
		}
	}
}

func TestSoma(t *testing.T) {
	test := sum(30,20)
	result := 50

	if test != result {
		t.Error("Expected: ", result, "but got: ", test)
	}
}

func TestSoma2(t *testing.T) {
	test := sum(70, 100)
	result := 170

	if test != result {
		t.Error("Expected: ", result, "but got: ", test)
	}
}

func TestSoma3(t *testing.T) {
	test := sum(70, 100)
	result := 170

	if test != result {
		t.Error("Expected: ", result, "but got: ", test)
	}
}