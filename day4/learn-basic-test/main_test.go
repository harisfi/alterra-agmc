package main

import (
	"io/ioutil"
	"os"
	"testing"
)

type testCase struct {
	arg1, arg2, expectation int
}

func TestMul(t *testing.T) {
	testCases := []testCase{
		{2, 3, 6},
		{3, 1, 3},
	}

	for _, testCase := range testCases {
		res := Mul(testCase.arg1, testCase.arg2)
		if res != testCase.expectation {
			t.Errorf("Expected %d, but got %d", testCase.expectation, res)
		}
	}
}

func TestDiv(t *testing.T) {
	testCases := []testCase{
		{6, 3, 2},
		{9, 3, 3},
	}

	for _, testCase := range testCases {
		res := Div(testCase.arg1, testCase.arg2)
		if res != testCase.expectation {
			t.Errorf("Expected %d, but got %d", testCase.expectation, res)
		}
	}
}

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "OK\n" {
		t.Errorf("Expected %s, got %s", "OK", out)
	}
}
