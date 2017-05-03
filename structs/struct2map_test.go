package structs

import "testing"

type Example struct {
	A string `json:"smaA" need:"yes"`
	B string `json:"bigB" need:"no"`
	C AAA    `json:"AAAC", need:"yes"`
}

type AAA struct {
	C bool `json:"booC" need:"yes"`
}

func TestStructToMap(t *testing.T) {
	a := &Example{
		A: "777",
		B: "999",
		C: AAA{C: true},
	}

	b := &Example{
		B: "787",
	}

	t.Logf("Example: %v", StructToMap(a))
	t.Logf("Example: %v", StructToMap(b))
}
