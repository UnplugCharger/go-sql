// file: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

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
	///**********************Querrying Multiple Birds*********************88//
	rows, err := db.Query("SELECT bird, description FROM birds limit 10")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}
	// create a slice of birds to hold our results
	birds := []Bird{}

	// iterate over the returned rows
	// we can go over to the next row by calling the `Next` method, which will
	// return `false` if there are no more rows
	for rows.Next() {
		bird := Bird{}
		// create an instance of `Bird` and write the result of the current row into it
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		// append the current instance to the slice of birds
		birds = append(birds, bird)
	}
	// print the length, and all the birds
	fmt.Printf("found %d birds: %+v", len(birds), birds)

	//***************Inserting data *****************************************************************************/////

	// sample data that we want to insert
	newBird := Bird{
		Species:     "Pelicano",
		Description: "Looong legs",
	}
	// the `Exec` method returns a `Result` type instead of a `Row`
	// we follow the same argument pattern to add query params
	result, err := db.Exec("INSERT INTO birds (bird, description) VALUES ($1, $2)", newBird.Species, newBird.Description)
	if err != nil {
		log.Fatalf("could not insert row: %v", err)
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	// we can log how many rows were inserted
	fmt.Println("inserted", rowsAffected, "rows")

}
