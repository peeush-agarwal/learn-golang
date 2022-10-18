package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "A",
		Price: 1.0,
		SKU:   "abc-def-ghi",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
