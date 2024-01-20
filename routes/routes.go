package routes

import (
	c "go_projects/productCatalog/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.NewProduct)
	http.HandleFunc("/create", c.CreateProduct)
	http.HandleFunc("/delete", c.DeleteProduct)
	http.HandleFunc("/edit", c.EditProduct)
	http.HandleFunc("/update", c.UpdateProduct)
}
