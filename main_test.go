package main

import (
	"testing"
)

type familyBusTest struct {
	output, families int
	members          string
}

var familyBusTests = []familyBusTest{
	{4, 5, "1 2 4 3 3"},
	{5, 8, "2 3 4 4 2 1 3 1"},
	{3, 5, "9 2 8"},
	{3, 6, "9 8 7"},
	{5, 0, "2 3"},
	// {2, 0, "a b"},
}

func TestFamilyBus(t *testing.T) {
	for i := 0; i < 100000; i++ {
		for _, test := range familyBusTests {
			output, err := familyBus(test.families, test.members)
			if err != "" {
				errT := "input must be equal with count of family"
				if err != errT {
					t.Errorf("Output %q not equal to expected %q", err, errT)
				}
			} else if output != test.output {
				t.Errorf("Output %q not equal to expected %q", output, test.output)
			}
		}
	}
}
