package models

import "testing"

func TestProductModel(t *testing.T) {
	p := &Product{}
	p.Name = "Apple Iphone 11"
	p.Price = 1000.00
	p.Quality = 100
	p.Status = ProductStatus_Available
	if err := p.Validate(); err != nil {
		t.Error(err)
	}
}
