package structs

import "testing"

type Example struct {
	A string `json:"smaA" hidden:"yes"`
	B string `json:"bigB" hidden:"no"`
}

func TestStructToMap(t *testing.T) {
	a := &Example{
		A: "777",
		B: "999",
	}

	b := &Example{
		B: "787",
	}

	t.Logf("Example: %v", StructToMap(a))
	t.Logf("Example: %v", StructToMap(b))
}
