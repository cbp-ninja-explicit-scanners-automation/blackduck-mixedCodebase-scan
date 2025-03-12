package data

import (
	"encoding/json"
	"fmt"
	"io"
)

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}
func (p Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(p)
}
func getNextID() int {
	lp := GetProducts()
	lastP := products[len(lp)-1]
	return lastP.Id + 1
}

var ErrProductNotFound error = fmt.Errorf("Product not found")

func findProduct(id int) (Product, int, error) {
	for pos, p := range products {
		if p.Id == id {
			return p, pos, nil
		}
	}
	return Product{}, -1, ErrProductNotFound
}

// Non compliant code
func isPrefixOf(xs, ys []int) bool {
	for i := 0; i < len(xs); i++ {
		if len(ys) == 0 || xs[i] != ys[i] {
			return false
		}
	}
	return true
}
