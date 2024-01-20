package main

import (
	"fmt"
	r "go_projects/productCatalog/routes"
	"net/http"
)

func main() {
	r.LoadRoutes()
	fmt.Println("Listen and serve on 8000 port...")
	http.ListenAndServe(":8000", nil)
}
