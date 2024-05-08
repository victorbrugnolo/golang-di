package main

import (
	"database/sql"
	"fmt"
	"github.com/victorbrugnolo/golang-di/product"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	usecase := product.NewProductRepository(db)

	prod, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(prod.Name)
}
