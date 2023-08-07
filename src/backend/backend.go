package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id        int
	Name      string
	Inventory int
	Price     int
}

type App struct {
	DB   *sql.DB
	Port string
}

func (app *App) Init() {
	app.DB, _ = sql.Open("sqlite3", "../../practiceit.db")
}

func (p *Product) Scan(db *sql.DB) {
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

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func (app *App) Run() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on port ", app.Port)
	log.Fatal(http.ListenAndServe(app.Port, nil))
}
