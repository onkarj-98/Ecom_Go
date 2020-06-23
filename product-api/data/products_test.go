package data

import "testing"

func TestCheckValidation(t * testing.T){
	p := &Product{
		Name : "Onakar",
		Price : 1.00,
		SKU: "abs-abc-def",
	}
	err := p.Validate()
	if err != nil{
		t.Fatal(err) 
	}
}