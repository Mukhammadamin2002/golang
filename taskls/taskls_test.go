package main

import (
	"testing"
)

func testConverter(t *testing.T) {
	var test = Task{"", "Housework", 5.00, "hour"}
	var result = Task{"Monday", "Housework", 4.00, "hour"}

	if test != result {
		t.Errorf("result %v test %v", result, test)
	}
}
