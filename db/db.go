package db

//_ sinaliza que a lib só será utilzada em tempo de execução
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	conn := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	createCatalogTable(db)

	return db
}

func createCatalogTable(db *sql.DB) {
	productCatalogs, err := db.Prepare("create table products_catalogs (id serial primary key, name varchar, description varchar, price decimal,  amount integer)")
	if err != nil {
		panic(err.Error())
	}

	productCatalogs.Exec()
}
