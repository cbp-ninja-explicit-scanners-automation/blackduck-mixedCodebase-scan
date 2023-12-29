package data

import (
	"html/template"
	"reflect"
	"time"
)

type Products []Product

// database mocking
var products = Products{
	{
		Id:          1,
		Name:        "Latte",
		Description: "Made of milk and light coffee",
		Price:       12.50,
		SKU:         "abc43",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		ImagePath:   template.URL("images/latte.jpg"),
	},
	{
		Id:          2,
		Name:        "IceTea",
		Description: "Tea with Ice Refreshing!!",
		Price:       28.30,
		SKU:         "hgt56",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		ImagePath:   template.URL("images/icetea.jpg"),
	},
}

// dto for product
type Product struct {
	Id          int          `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Price       float32      `json:"price,omitempty"`
	SKU         string       `json:"sku,omitempty"`
	ImagePath   template.URL `json:"image_path"`
	CreatedAt   string       `json:"created_at,omitempty"`
	UpdatedAt   string       `json:"updated_at,omitempty"`
	DeletedAt   string       `json:"-"`
}

func GetProducts() Products {
	return products
}

func GetProduct(id int) (Product, error) {
	p, _, err := findProduct(id)
	if err != nil {
		return Product{}, err
	}
	return p, nil

}

func AddProduct(p Product) {
	p.Id = getNextID()
	products = append(products, p)

}

func UpdateProduct(p Product, id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	productList := GetProducts()

	old := productList[pos]
	new := p
	finalProd := Product{}

	oldProd := reflect.ValueOf(old)
	newProd := reflect.ValueOf(new)
	final := reflect.ValueOf(finalProd).Elem()

	for i := 0; i < oldProd.NumField(); i++ {
		switch newProd.Field(i).Kind() {
		case reflect.Int, reflect.Float32:
			if newProd.Field(i).IsZero() {
				final.Field(i).Set(oldProd.Field(i))
			} else {
				final.Field(i).Set(newProd.Field(i))
			}
		case reflect.String:
			if newProd.Field(i).String() == "" {
				final.Field(i).Set(oldProd.Field(i))
			} else {
				final.Field(i).Set(newProd.Field(i))
			}
		}
	}
	finalProd.UpdatedAt = time.Now().UTC().String()
	productList[pos] = finalProd

	return nil
}
