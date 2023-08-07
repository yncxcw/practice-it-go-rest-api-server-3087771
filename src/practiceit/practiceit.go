package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id        int
	Name      string
	Inventory int
	Price     int
}

func main() {
	// backend.Run(":9003")
	fmt.Println("Opening table")
	db, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Querying tables")
	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	defer rows.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var p Product
		rows.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price)
		println(p.Id, p.Name, p.Inventory, p.Price)
	}
}
