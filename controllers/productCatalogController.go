package controllers

import (
	m "go_projects/productCatalog/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//carrega os templates da pasta "templates"
var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// sempre que o path for "/" a função "index" será chamada pelo "HandleFunc".
	// Ela deve receber a interface ResponseWriter, responsável por construir a resposta,
	// e um ponteiro para a requisição em questão.

	templates.ExecuteTemplate(w, "index", m.ListProducts())
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "new", nil)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// pegando o valor inserido no formulário, de acordo com o nome que foi dado ao campo(create.html).
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		// O valor que vem do formulário é sempre uma string, precisamos fazer um parse para um tipo comum do go.
		parsedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Falha ao converter o preço", err)
			http.Redirect(w, r, "/", 301)
			return
		}

		parsedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Falha ao converter a quantidade", err)
			http.Redirect(w, r, "/", 301)
			return
		}

		m.CreateProduct(name, description, parsedPrice, parsedAmount)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//pegando o id através da URL
	productID := r.URL.Query().Get("id")
	m.DeleteProduct(productID)

	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := m.FetchProduct(productID)
	templates.ExecuteTemplate(w, "edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		parsedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Falha ao converter o id", err)
			http.Redirect(w, r, "/", 301)
			return
		}

		parsedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Falha ao converter o preço", err)
			http.Redirect(w, r, "/", 301)
			return
		}

		parsedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Falha ao converter a quantidade", err)
			http.Redirect(w, r, "/", 301)
			return
		}

		m.UpdateProduct(parsedId, name, description, parsedPrice, parsedAmount)
	}
	http.Redirect(w, r, "/", 301)
}
