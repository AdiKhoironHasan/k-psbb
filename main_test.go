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
}

func TestFamilyBus(t *testing.T) {
	for i := 0; i < 100000; i++ {
		for _, test := range familyBusTests {
			if output, _ := familyBus(test.families, test.members); output != test.output {
				t.Errorf("Output %q not equal to expected %q", output, test.output)
			}
		}
	}
}
