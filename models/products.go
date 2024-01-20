package models

import (
	"database/sql"
	"go_projects/productCatalog/db"
)

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func ListProducts() []Product {
	db := db.DbConnect()
	//define algo que será executado por ultimo, utilize para garantir a limpeza dos recursos.
	defer db.Close()

	allProducts, err := db.Query("select * from products_catalogs")
	if err != nil {
		panic(err.Error())
	}

	return scanProducts(allProducts)
}

func scanProducts(allProducts *sql.Rows) []Product {
	//Utiliza as linhas retornadas na query para criar um slice. A criação desse slice segue os passos:
	// 1 - A função Next() prepara as linhas para que a função Scan() consiga ler;
	// 2 - Os tipos equivalentes a estrutura que pretendemos retornar devem ser definidos;
	// 3 - A função Scan() converte cada linha retornada pela query em um tipo comum no go. Ela coloca o campo que veio do bando direto na memória da variável;
	// 4 - Os valores de cada tipo que compõe a struct são definidos com base na definição do passo 1;
	// 5 - O slice "products" será acrescentado com structs(p).

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := allProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Id = id
		p.Amount = amount
		p.Price = price
		p.Description = description

		products = append(products, p)
	}

	return products
}

func FetchProduct(id string) Product {
	db := db.DbConnect()
	defer db.Close()

	//faz a query com parametro direto, sem preparar antes
	productDb, err := db.Query("select * from products_catalogs where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	return scanProduct(productDb)
}

func scanProduct(productDb *sql.Rows) Product {
	product := Product{}

	for productDb.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := productDb.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Id = id
		product.Amount = amount
		product.Price = price
		product.Description = description
	}

	return product
}

func CreateProduct(name string, description string, price float64, amount int) {
	db := db.DbConnect()
	defer db.Close()

	//preparando a inserção para ser executada mais adiante
	insertDbData, err := db.Prepare("insert into products_catalogs(name, description, price, amount) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	//executando a inserção que foi preparada anteriormente
	insertDbData.Exec(name, description, price, amount)
}

func DeleteProduct(id string) {
	db := db.DbConnect()
	defer db.Close()
	deleteDbData, err := db.Prepare("delete from products_catalogs where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteDbData.Exec(id)
}

func UpdateProduct(id int, name string, description string, price float64, amount int) {
	db := db.DbConnect()
	defer db.Close()

	updateDbData, err := db.Prepare("update products_catalogs set name = $1, description = $2, price = $3, amount = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateDbData.Exec(name, description, price, amount, id)
}
