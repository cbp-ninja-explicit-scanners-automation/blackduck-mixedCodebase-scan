package handlers

import (
	"fmt"
	"html/template"
	"main/zero1/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	BasePath   = "/Users/ext.amit.r/Documents/projectZero/zero1/src/templates"
	Base       = fmt.Sprintf("%s/base.html", BasePath)
	Index      = fmt.Sprintf("%s/index.html", BasePath)
	Footer     = fmt.Sprintf("%s/footer.html", BasePath)
	About      = fmt.Sprintf("%s/about.html", BasePath)
	AddProduct = fmt.Sprintf("%s/product.html", BasePath)
)

func (p *Product) GetTemplProduct(rw http.ResponseWriter, r *http.Request) {

	// getting product id from url
	vars := mux.Vars(r)
	id := vars["id"]

	prodId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(prodId)

	//templating
	tmpl := template.Must(template.ParseFiles(Base, Index, Footer))
	product, err := data.GetProduct(prodId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	tmpl.ExecuteTemplate(rw, "index.html", product)

}

func (p *Product) GetTemplProducts(rw http.ResponseWriter, r *http.Request) {
	//templating
	tmpl := template.Must(template.ParseFiles(Base, Index, Footer))
	products := data.GetProducts()
	tmpl.ExecuteTemplate(rw, "index.html", products)
}

func (p *Product) GetTemplAddProduct(rw http.ResponseWriter, r *http.Request) {
	//templating
	tmpl := template.Must(template.ParseFiles(Base, AddProduct, Footer))
	tmpl.ExecuteTemplate(rw, "product.html", nil)
}

func (p *Product) PostTemplAddProduct(rw http.ResponseWriter, r *http.Request) {
	//templating
	tmpl := template.Must(template.ParseFiles(Base, AddProduct, Footer))
	price, err := strconv.ParseFloat(r.FormValue("productPrice"), 32)
	if err != nil {
		price = 0
	}
	product := data.Product{
		Name:        r.FormValue("productName"),
		Description: r.FormValue("productDesc"),
		Price:       float32(price),
		SKU:         r.FormValue("productSKU"),
		ImagePath:   template.URL("images/defaultCoffee.jpg"),
	}
	data.AddProduct(product)
	tmpl.ExecuteTemplate(rw, "product.html", nil)
}
