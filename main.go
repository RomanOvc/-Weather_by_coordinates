package main

import (
	"WeatherByCoordinates/api"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "acer5800"
	dbname   = "lol1"
)

func main() {

	res1, err := api.FullResult("New York")
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
