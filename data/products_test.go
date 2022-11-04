package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "rsuky",
		Price: 1.00,
		SKU:   "absd-xsrt-kqzs",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
