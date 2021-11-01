// file: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Bird struct {
	Species     string
	Description string
}








func main() {
	// The `sql.Open` function opens a new `*sql.DB` instance. We specify the driver name
	// DSN string
	db, err := sql.Open("pgx", "user=postgres password=new_password host=localhost port=5432 database=students sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")



	// `QueryRow` always returns a single row from the database
row := db.QueryRow("SELECT bird, description FROM birds LIMIT 1")
// Create a new `Bird` instance to hold our query results
bird := Bird{}
// the retrieved columns in our row are written to the provided addresses
// the arguments should be in the same order as the columns defined in
// our query
if err := row.Scan(&bird.Species, &bird.Description); err != nil {
	log.Fatalf("could not scan row: %v", err)
}
fmt.Printf("found bird: %+v\n", bird)




}