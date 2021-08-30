package main

import (
	"WeatherByCoordinates/api"
	"WeatherByCoordinates/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "acer5800"
	dbname   = "lol1"
)

func initPostgresInstance() (*sql.DB, error) {
	// create db conn object with params
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=%s",
		host,
		user,
		dbname,
		password,
		port,
		"false",
	))
	if err != nil {
		return nil, errors.Wrap(err, "postgres sql open connect")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "postgres ping error")
	}

	log.Printf("Postgres connected on %d port", port)
	return db, nil
}

func main() {

	db, err := initPostgresInstance()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.Repository{Db: db}
	usecase := api.UseCase{Repo: repo}

	res1, err := usecase.FullResult("New York")
	if err != nil {
		log.Print(err)
	}

	u, err := json.Marshal(res1)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(u))
	// connection string
	// connection string
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlconn)
	// CheckError(err)

	// defer db.Close()

	// // insert
	// // hardcoded
	// insertStmt := `insert into students ("name", "role") values('Johna', 1)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// // dynamic
	// insertDynStmt := `insert into students("name", "role") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jaane", 2)
	// CheckError(e)

	// insertDynStmt1 := `insert into students("name", "role") values($1, $2)`
	// _, e = db.Exec(insertDynStmt1, "Janes", 2)
	// CheckError(e)

}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
