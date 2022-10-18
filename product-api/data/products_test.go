package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "A",
		Price: 1.0,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
